package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/Sterl1337/job-hunt-ai/dashboard/internal/data"
	"github.com/Sterl1337/job-hunt-ai/dashboard/internal/theme"
	"github.com/Sterl1337/job-hunt-ai/dashboard/internal/ui/screens"
)

// dataReloadMsg is sent when applications.md changes on disk.
type dataReloadMsg struct{ modTime time.Time }

// fileCheckMsg is sent when the periodic check finds no changes.
type fileCheckMsg struct{ modTime time.Time }

type viewState int

const (
	viewPipeline viewState = iota
	viewReport
)

type appModel struct {
	pipeline      screens.PipelineModel
	viewer        screens.ViewerModel
	state         viewState
	careerOpsPath string
	lastModTime   time.Time
}

// appsFilePath returns the path to applications.md, checking both locations.
func appsFilePath(careerOpsPath string) string {
	p := filepath.Join(careerOpsPath, "data", "applications.md")
	if _, err := os.Stat(p); err == nil {
		return p
	}
	return filepath.Join(careerOpsPath, "applications.md")
}

// watchApps polls applications.md every 2s and sends a reload message if it changed.
func watchApps(careerOpsPath string, lastMod time.Time) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(2 * time.Second)
		p := appsFilePath(careerOpsPath)
		info, err := os.Stat(p)
		if err != nil {
			return fileCheckMsg{modTime: lastMod}
		}
		if info.ModTime().After(lastMod) {
			return dataReloadMsg{modTime: info.ModTime()}
		}
		return fileCheckMsg{modTime: info.ModTime()}
	}
}

func (m appModel) Init() tea.Cmd {
	return watchApps(m.careerOpsPath, m.lastModTime)
}

func (m appModel) reloadApps() appModel {
	apps := data.ParseApplications(m.careerOpsPath)
	metrics := data.ComputeMetrics(apps)
	old := m.pipeline
	t := theme.NewTheme("catppuccin-mocha")
	m.pipeline = screens.NewPipelineModel(t, apps, metrics, m.careerOpsPath, old.Width(), old.Height())
	m.pipeline.CopyReportCache(&old)
	return m
}

func (m appModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.pipeline.Resize(msg.Width, msg.Height)
		if m.state == viewReport {
			m.viewer.Resize(msg.Width, msg.Height)
		}
		pm, cmd := m.pipeline.Update(msg)
		m.pipeline = pm
		return m, cmd

	case dataReloadMsg:
		m = m.reloadApps()
		now := msg.modTime
		m.lastModTime = now
		m.pipeline.SetLastReloaded(now)
		return m, watchApps(m.careerOpsPath, now)

	case fileCheckMsg:
		m.lastModTime = msg.modTime
		return m, watchApps(m.careerOpsPath, msg.modTime)

	case screens.PipelineRequestReloadMsg:
		// "r" key: force reload regardless of mod time
		p := appsFilePath(m.careerOpsPath)
		info, _ := os.Stat(p)
		m = m.reloadApps()
		var modTime time.Time
		if info != nil {
			modTime = info.ModTime()
		}
		m.lastModTime = modTime
		m.pipeline.SetLastReloaded(time.Now())
		return m, watchApps(m.careerOpsPath, modTime)

	case screens.PipelineClosedMsg:
		return m, tea.Quit

	case screens.PipelineLoadReportMsg:
		archetype, tldr, remote, comp := data.LoadReportSummary(msg.CareerOpsPath, msg.ReportPath)
		m.pipeline.EnrichReport(msg.ReportPath, archetype, tldr, remote, comp)
		return m, nil

	case screens.PipelineUpdateStatusMsg:
		err := data.UpdateApplicationStatus(msg.CareerOpsPath, msg.App, msg.NewStatus)
		if err != nil {
			return m, nil
		}
		apps := data.ParseApplications(m.careerOpsPath)
		metrics := data.ComputeMetrics(apps)
		old := m.pipeline
		m.pipeline = screens.NewPipelineModel(
			theme.NewTheme("catppuccin-mocha"),
			apps, metrics, m.careerOpsPath,
			old.Width(), old.Height(),
		)
		m.pipeline.CopyReportCache(&old)
		return m, nil

	case screens.PipelineOpenReportMsg:
		m.viewer = screens.NewViewerModel(
			theme.NewTheme("catppuccin-mocha"),
			msg.Path, msg.Title,
			m.pipeline.Width(), m.pipeline.Height(),
		)
		m.state = viewReport
		return m, nil

	case screens.ViewerClosedMsg:
		m.state = viewPipeline
		return m, nil

	case screens.PipelineOpenURLMsg:
		url := msg.URL
		return m, func() tea.Msg {
			var cmd *exec.Cmd
			switch runtime.GOOS {
			case "darwin":
				cmd = exec.Command("open", url)
			case "linux":
				cmd = exec.Command("xdg-open", url)
			case "windows":
				cmd = exec.Command("cmd", "/c", "start", "", url)
			default:
				cmd = exec.Command("xdg-open", url)
			}
			_ = cmd.Start()
			return nil
		}

	default:
		if m.state == viewReport {
			vm, cmd := m.viewer.Update(msg)
			m.viewer = vm
			return m, cmd
		}
		pm, cmd := m.pipeline.Update(msg)
		m.pipeline = pm
		return m, cmd
	}
}

func (m appModel) View() string {
	if m.state == viewReport {
		return m.viewer.View()
	}
	return m.pipeline.View()
}

func main() {
	pathFlag := flag.String("path", ".", "Path to career-ops directory")
	flag.Parse()

	careerOpsPath := *pathFlag

	// Load applications
	apps := data.ParseApplications(careerOpsPath)
	if apps == nil {
		fmt.Fprintf(os.Stderr, "Error: could not find applications.md in %s or %s/data/\n", careerOpsPath, careerOpsPath)
		os.Exit(1)
	}

	// Get initial mod time for file watcher
	var initModTime time.Time
	if info, err := os.Stat(appsFilePath(careerOpsPath)); err == nil {
		initModTime = info.ModTime()
	}

	// Compute metrics
	metrics := data.ComputeMetrics(apps)

	// Build pipeline model
	t := theme.NewTheme("catppuccin-mocha")
	pm := screens.NewPipelineModel(t, apps, metrics, careerOpsPath, 120, 40)

	// Batch-load all report summaries
	for _, app := range apps {
		if app.ReportPath == "" {
			continue
		}
		archetype, tldr, remote, comp := data.LoadReportSummary(careerOpsPath, app.ReportPath)
		if archetype != "" || tldr != "" || remote != "" || comp != "" {
			pm.EnrichReport(app.ReportPath, archetype, tldr, remote, comp)
		}
	}

	pm.SetLastReloaded(initModTime)

	m := appModel{
		pipeline:      pm,
		careerOpsPath: careerOpsPath,
		lastModTime:   initModTime,
	}

	p := tea.NewProgram(m, tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
