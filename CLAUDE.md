# job-hunt-ai — Claude guidance

## Origin

This repository is maintained by [Sterling Fisher](https://github.com/Sterl1337) as an adapted version of the original upstream project [santifer/career-ops](https://github.com/santifer/career-ops).

Sterling customized and maintains this repo, but he is not the original upstream author. Do not describe Sterling as the original creator of the upstream project. Preserve attribution to the original source and describe this repository as Sterling's maintained adaptation.

## Project overview

job-hunt-ai is a terminal-first job-search workflow built around role evaluation, application tracking, resume customization, portal scanning, and dashboard visibility.

## Main commands

- `/job-hunt-ai {job URL or JD}` — run the main evaluation pipeline
- `/job-hunt-ai scan` — scan configured portals
- `/job-hunt-ai pdf` — generate resume output
- `/job-hunt-ai tracker` — review application status
- `/job-hunt-ai batch` — process multiple opportunities
- `/job-hunt-ai pipeline` — process pending URLs
- `/job-hunt-ai apply` — prepare application data without auto-submitting
- `/job-hunt-ai mass-apply` — multi-platform mass-apply workflow

## Working rules

- Treat `cv.md`, `config/profile.yml`, `modes/_profile.md`, `article-digest.md`, and `portals.yml` as user-specific files.
- Keep personal data out of system files.
- Use defensive file checks before reading or writing files.
- Prefer updating user-layer files for personalization.
- Do not auto-submit job applications.
- Favor quality over volume when evaluating roles.

## Repo conventions

- Node.js scripts use `.mjs`
- Dashboard code lives in `dashboard/`
- Reports live in `reports/`
- Output files live in `output/`
- Configuration is primarily Markdown and YAML

## Notes

If the user wants to customize archetypes, scoring, templates, or workflows, edit the relevant repo files directly and keep attribution language accurate.
