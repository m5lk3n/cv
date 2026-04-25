# Kanban Board

## Backlog

- [ ] support additional themes?
- [ ] favicon.png should be the resume.json's basic.image (use app favicon as fallback)
- [ ] check out https://github.com/igolaizola/webcli

## Open

- [ ] add configuration?
  - [ ] background color
  - [ ] foreground color
  - [ ] links
  - [ ] site link/title

## To do

- [ ] entry page to choose between static and dynamic
- [ ] fix links incl. hardcodings like https://registry.jsonresume.org/m5lk3n in `index.html`
- [ ] add `.env.example`
- [ ] improve wording
- [ ] complete `README.md`, also include screenshots
  - [ ] add `screenshots/animated.gif`

## Doing

- [ ] prep repo https://github.com/m5lk3n/cv
- [ ] fix package name

### Done

- [x] add favicon
- [x] add link to GitHub repo
- [x] support clickable URLs
- [x] commands
  - [x] version
    - [x] embed semver in binary, read version from there
  - [x] about / summary
  - [x] experience / work
    - [x] achievements
  - [x] register `achievement(s)` as top-level command for `work achievement(s)`
  - [x] skills
  - [x] certification(s) (license(s) as alias)
  - [x] basics
    - [x] pages / links / online / social
  - [x] education
  - [x] language(s)
  - [x] volunteering
  - [x] publications
  - [x] references (recommendations)
  - [x] faq
  - [x] projects
- [x] address FIXMEs
- [x] add licenses
- [x] https://cv.michael-klein.info/
- [x] resume as downloadable PDF
- [x] add link to published resume under registry.jsonresume.org
- [x] add `publish-to-web` target to Makefile
- [x] publish `resume.json` under https://registry.jsonresume.org/
  - [x] using `Makefile`
  - [ ] ~~using GitHub workflow (see https://github.com/thomasdavis/resume)~~
