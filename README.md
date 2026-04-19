# job-hunt-ai

> AI-powered job search pipeline built on Claude Code. Evaluate offers, generate tailored resumes, scan portals, run application workflows, and track everything from your terminal.

[![Version](https://img.shields.io/badge/version-1.2.0-blue.svg)](./VERSION)
[![Claude Code](https://img.shields.io/badge/Claude-Code-7C3AED.svg)](https://claude.ai/code)
[![Node](https://img.shields.io/badge/node-%3E%3D20-green.svg)](https://nodejs.org/)
[![License](https://img.shields.io/badge/license-MIT-lightgrey.svg)](./LICENSE)

---

## Overview & Attribution

job-hunt-ai is maintained by **Sterling Fisher** as a customized adaptation of the original open-source project **[career-ops](https://github.com/santifer/career-ops)** by **santifer**.

This repository reflects Sterling’s branding, workflow changes, documentation updates, and automation experiments, but it is **not** presented as the original upstream project. Preserve credit to the upstream source when describing the repo’s origin.

---

## What it does

- Evaluates job descriptions and scores fit
- Generates ATS-friendly resumes and PDFs
- Tracks applications and pipeline state
- Scans job portals for matching roles
- Supports live apply assistance
- Adds **mass-apply orchestration** across multiple platforms

---

## New in v1.2.0: Mass Apply

Mass apply adds a conductor-style workflow for higher-volume campaigns across:

- LinkedIn Easy Apply
- ZipRecruiter 1-Click / Quick Apply
- Glassdoor Easy Apply
- Indeed Easily Apply

Key behavior:

- platform-specific apply modes
- deduping against tracker/logs before applying
- pause-before-submit by default
- optional turbo / auto-submit behavior when explicitly enabled by the user
- session summaries and tracker integration after each run

Primary command:

```bash
/job-hunt-ai mass-apply
```

The router and skill manifest recognize the mass-apply command, and the repo includes platform mode files for LinkedIn, ZipRecruiter, and Glassdoor.

---

## Core commands

```bash
/job-hunt-ai {JD}        # auto-pipeline from pasted JD or JD URL
/job-hunt-ai pipeline    # process pending URLs from inbox
/job-hunt-ai oferta      # evaluation only
/job-hunt-ai ofertas     # compare multiple roles
/job-hunt-ai contacto    # contact and outreach support
/job-hunt-ai deep        # deep company research prompt
/job-hunt-ai pdf         # generate ATS resume PDF
/job-hunt-ai tracker     # view application status
/job-hunt-ai apply       # live application assistant
/job-hunt-ai scan        # scan portals for roles
/job-hunt-ai batch       # batch processing
/job-hunt-ai mass-apply  # multi-platform mass application campaign
```

---

## Mass apply quick start

```bash
npm install
npx playwright install chromium
claude --chrome --dangerously-skip-permissions
```

Then, inside the Claude-controlled Chrome session:

1. Install the Claude browser extension if you have not already.
2. Log into each supported board you want to use for the campaign:
   - LinkedIn
   - ZipRecruiter
   - Glassdoor
   - Indeed
3. Confirm your resumes, profile answers, and saved defaults are current.
4. Run:

```bash
/job-hunt-ai mass-apply
```

Recommended run pattern:
- start in pause-before-submit mode
- verify filters, resume selection, and question handling
- only enable faster or more automated submission behavior after validation

---

## Files added or updated for mass apply

- `.claude/skills/career-ops/SKILL.md` (skill router branded to invoke `/job-hunt-ai`)
- `CLAUDE.md`
- `modes/mass-apply.md`
- `modes/linkedin-apply.md`
- `modes/ziprecruiter-apply.md`
- `modes/glassdoor-apply.md`
- supporting docs and batch workflow files from the update set

---

## Quick start

```bash
npm install
npm run doctor
```

Then run the command you need, such as:

```bash
/job-hunt-ai mass-apply
```

---

## Credits & Attribution

- **Original upstream project:** [santifer/career-ops](https://github.com/santifer/career-ops)
- **Maintained adaptation in this repo:** [Sterl1337/job-hunt-ai](https://github.com/Sterl1337/job-hunt-ai)
- **Current maintainer and customization work:** Sterling Fisher

This fork includes Sterling’s repo-specific customization work, documentation updates, workflow changes, and additional automation experiments layered on top of the upstream foundation.

## License

MIT
