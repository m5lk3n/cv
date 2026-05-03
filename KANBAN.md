# Kanban Board

## To do

- [ ] try to improve UX on smaller screens like on a smartphone
- [ ] consolidate API entities and structs from `resume/resume.go` into a single model
- [ ] animated pointer to PDF download?
- [ ] `favicon.png` could be the `resume.json`'s `basic.image` (use app favicon as fallback)
- [ ] make it bubbletea TUI?
- [ ] check out https://github.com/igolaizola/webcli

## Doing

- [ ] add `publish-api` target to Makefile, using via `deploy/` scripting
- [ ] containerize `cvapi`

### Done

- [x] add favicon
- [x] add link to GitHub repo
- [x] support clickable URLs
- [x] commands
  - [x] `version`
    - [x] embed semver in binary, read version from there
  - [x] `about` / summary
  - [x] `experience` / work
    - [x] `achievements`
  - [x] register `achievements` as top-level command for `work achievements`
  - [x] `skills`
  - [x] `certifications` (`license(s)` as alias)
  - [x] `basics`
    - [x] pages / links / online / social
  - [x] `education`
  - [x] `languages`
  - [x] `volunteering`
  - [x] `publications`
  - [x] `references` (recommendations)
  - [x] `faqs`
  - [x] `projects`
- [x] address FIXMEs
- [x] add licenses
- [x] https://cv.michael-klein.info/
- [x] resume as downloadable PDF
- [x] publish `resume.json` under https://registry.jsonresume.org/
- [x] add link to published resume under https://registry.jsonresume.org/
- [x] add `publish-to-web` target to Makefile
  - [x] using `Makefile`
  - [ ] ~~using GitHub workflow (see https://github.com/thomasdavis/resume)~~
- [x] move repo to https://github.com/m5lk3n/cv
- [x] fix hardcodings
- [x] add `.env.example`
- [x] complete `README.md`
  - [x] add screenshot(s)
- [x] improve CV wording
- [x] add splash screen to choose between static and interactive
- [x] add tests
- [x] support additional themes
- [x] embed app version in splash screen
- [x] add API
