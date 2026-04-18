# job-hunt-ai

> AI-assisted job search workflow for evaluating roles, tailoring resumes, scanning portals, and tracking applications from the terminal.

![Claude Code](https://img.shields.io/badge/Claude_Code-000?style=flat&logo=anthropic&logoColor=white)
![Node.js](https://img.shields.io/badge/Node.js-339933?style=flat&logo=node.js&logoColor=white)
![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)
![Playwright](https://img.shields.io/badge/Playwright-2EAD33?style=flat&logo=playwright&logoColor=white)
![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)

## Overview

job-hunt-ai is a practical job-search automation workflow built around role evaluation, application tracking, resume customization, and terminal-first execution.

This repository is maintained by **Sterling Fisher** and represents his adapted version of the original open-source project **[career-ops](https://github.com/santifer/career-ops)** by **santifer**. This version includes Sterling's customizations, workflow changes, branding, and ongoing improvements, but it is not presented as the original upstream project.

## What it does

- Evaluates opportunities with a structured scoring workflow
- Generates tailored resume outputs from a source CV
- Scans job portals and company career pages
- Tracks applications and reports in one place
- Supports batch processing and terminal-first workflows
- Provides a dashboard for pipeline visibility and review

## Why this project matters

This repository highlights practical work in:

- automation
- terminal tooling
- documentation
- workflow design
- job-search operations
- AI-assisted productivity systems

## Quick start

```bash
# Clone and install
git clone https://github.com/Sterl1337/job-hunt-ai.git
cd job-hunt-ai
npm install
npx playwright install chromium

# Configure
cp config/profile.example.yml config/profile.yml
cp templates/portals.example.yml portals.yml

# Add your CV source file
# Create cv.md in the project root

# Launch Claude Code
claude
```

See `docs/SETUP.md` for full setup details.

## Usage

```bash
/career-ops                # Show commands
/career-ops scan           # Scan portals
/career-ops pdf            # Generate tailored resume output
/career-ops tracker        # View application status
/career-ops batch          # Batch process opportunities
/career-ops pipeline       # Process pending URLs
```

You can also paste a job URL directly into the workflow and let the pipeline evaluate it.

## Project structure

```text
job-hunt-ai/
├── CLAUDE.md
├── config/
├── templates/
├── modes/
├── batch/
├── dashboard/
├── docs/
├── data/
├── reports/
└── output/
```

## Tech stack

- Claude Code
- Node.js
- Playwright
- Go (dashboard)
- Markdown / YAML / TSV workflow files

## Attribution

- **Original upstream project:** [santifer/career-ops](https://github.com/santifer/career-ops)
- **Maintained adaptation in this repo:** [Sterl1337/job-hunt-ai](https://github.com/Sterl1337/job-hunt-ai)
- **Current maintainer and customization work:** Sterling Fisher

## About this version

Sterling Fisher maintains this version as a customized adaptation focused on practical workflow improvements, personal automation, and portfolio-ready documentation.

## License

MIT
