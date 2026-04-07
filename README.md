# job-hunt-ai

> AI-powered job search pipeline built on Claude Code. Evaluate offers, generate tailored CVs, scan portals, and track everything — all from your terminal.

![Claude Code](https://img.shields.io/badge/Claude_Code-000?style=flat&logo=anthropic&logoColor=white)
![Node.js](https://img.shields.io/badge/Node.js-339933?style=flat&logo=node.js&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Playwright](https://img.shields.io/badge/Playwright-2EAD33?style=flat&logo=playwright&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)

---

## What It Does

job-hunt-ai turns Claude Code into a full job search command center. Instead of manually tracking applications in a spreadsheet, you get an AI pipeline that:

- **Evaluates offers** with a structured A-F scoring system (10 weighted dimensions)
- **Generates tailored PDFs** — ATS-optimized CVs customized per job description
- **Scans portals** automatically (Greenhouse, Ashby, Lever, company career pages)
- **Processes in batch** — evaluate 10+ offers in parallel with sub-agents
- **Tracks everything** in a single source of truth with integrity checks

**This is a quality filter, not a spray-and-pray tool.** The system recommends against applying to anything scoring below 4.0/5. Your time matters, and so does the recruiter's.

## Features

| Feature | Description |
|---------|-------------|
| **Auto-Pipeline** | Paste a URL, get a full evaluation + PDF + tracker entry |
| **6-Block Evaluation** | Role summary, CV match, level strategy, comp research, personalization, interview prep |
| **Interview Story Bank** | Builds STAR+R stories across evaluations — ready answers for any behavioral question |
| **ATS PDF Generation** | Keyword-injected CVs with clean typography |
| **Portal Scanner** | 45+ companies pre-configured + custom queries across Ashby, Greenhouse, Lever, Wellfound |
| **Batch Processing** | Parallel evaluation with `claude -p` workers |
| **Dashboard TUI** | Terminal UI with live auto-reload, pipeline funnel, score bars, filter tabs |
| **Human-in-the-Loop** | AI evaluates and recommends — you make every final call |
| **Pipeline Integrity** | Automated merge, dedup, status normalization, health checks |

## Quick Start

```bash
# 1. Clone and install
git clone https://github.com/Sterl1337/job-hunt-ai.git
cd job-hunt-ai && npm install
npx playwright install chromium   # Required for PDF generation

# 2. Configure
cp config/profile.example.yml config/profile.yml  # Fill in your details
cp templates/portals.example.yml portals.yml       # Customize target companies

# 3. Add your CV
# Create cv.md in the project root — this is the source of truth for all evaluations

# 4. Open Claude Code and let it onboard you
claude

# 5. Start using it
# Paste a job URL or run /career-ops
```

See [docs/SETUP.md](docs/SETUP.md) for the full setup guide.

## Usage

```
/career-ops                → Show all commands
/career-ops {paste a JD}   → Full auto-pipeline (evaluate + PDF + tracker)
/career-ops scan           → Scan portals for new offers
/career-ops pdf            → Generate ATS-optimized CV
/career-ops batch          → Batch evaluate multiple offers
/career-ops tracker        → View application status
/career-ops apply          → Fill application forms with AI
/career-ops pipeline       → Process pending URLs
/career-ops contacto       → LinkedIn outreach message
/career-ops deep           → Deep company research
/career-ops training       → Evaluate a course or cert
/career-ops project        → Evaluate a portfolio project
```

Or paste a job URL directly — career-ops detects it and runs the full pipeline automatically.

## How It Works

```
You paste a job URL or description
        │
        ▼
┌──────────────────┐
│  Archetype       │  Classifies role type from JD
│  Detection       │
└────────┬─────────┘
         │
┌────────▼──────────┐
│  A-F Evaluation   │  Match, gaps, comp research, STAR stories
│  (reads cv.md)    │
└────────┬──────────┘
         │
    ┌────┼────┐
    ▼    ▼    ▼
 Report  PDF  Tracker
  .md   .pdf   .tsv
```

## Dashboard TUI

The built-in terminal dashboard gives you a live view of your pipeline:

```bash
cd dashboard
go build -o career-dashboard.exe .
./career-dashboard.exe --path ..
```

Features: live auto-reload (2s polling), pipeline funnel bar, score mini-bars, 6 filter tabs, 4 sort modes, inline status changes, `r` to force reload.

## Pre-configured Portals

**AI Labs:** Anthropic, OpenAI, Mistral, Cohere, LangChain, Pinecone
**Voice AI:** ElevenLabs, PolyAI, Hume AI, Deepgram, Vapi, Bland AI
**Platforms:** Retool, Airtable, Vercel, Temporal, Glean, Arize AI
**Contact Center:** Ada, LivePerson, Sierra, Decagon, Talkdesk, Genesys
**Enterprise:** Salesforce, Twilio, Gong, Dialpad
**LLMOps:** Langfuse, Weights & Biases, Lindy, Cognigy
**Automation:** n8n, Zapier, Make.com

**Job boards:** Ashby, Greenhouse, Lever, Wellfound, Workable, RemoteFront

## Project Structure

```
job-hunt-ai/
├── CLAUDE.md                    # Agent instructions
├── cv.md                        # Your CV (create this, gitignored)
├── article-digest.md            # Your proof points (optional, gitignored)
├── config/
│   └── profile.example.yml      # Template for your profile
├── modes/                       # 14 skill modes
│   ├── _shared.md               # Shared context
│   ├── oferta.md                # Single offer evaluation
│   ├── pdf.md                   # PDF generation
│   ├── scan.md                  # Portal scanner
│   ├── batch.md                 # Batch processing
│   └── ...
├── templates/
│   ├── cv-template.html         # ATS-optimized CV template
│   ├── portals.example.yml      # Scanner config template
│   └── states.yml               # Canonical pipeline statuses
├── batch/
│   ├── batch-prompt.md          # Worker prompt template
│   └── batch-runner.sh          # Batch orchestrator
├── dashboard/                   # Go TUI pipeline viewer
├── data/                        # Your tracking data (gitignored)
├── reports/                     # Evaluation reports (gitignored)
├── output/                      # Generated PDFs (gitignored)
├── fonts/                       # Space Grotesk + DM Sans
└── docs/                        # Setup, customization, architecture
```

## Stack

- **Agent:** Claude Code with custom skill modes
- **PDF:** Playwright + HTML/CSS template
- **Scanner:** Playwright + Greenhouse API + WebSearch
- **Dashboard:** Go + Bubble Tea + Lipgloss (Catppuccin Mocha)
- **Data:** Markdown tables + YAML config + TSV batch files

## About

Built by Sterling Fisher — CompTIA Security+ / A+ certified, IT and automation specialist based in Atlanta, GA. Built this to run a smarter job search: fewer applications, better targeting, zero wasted time.

[![GitHub](https://img.shields.io/badge/GitHub-181717?style=for-the-badge&logo=github&logoColor=white)](https://github.com/Sterl1337)
[![Email](https://img.shields.io/badge/Email-EA4335?style=for-the-badge&logo=gmail&logoColor=white)](mailto:SterlingSFisher@gmail.com)

## License

MIT
