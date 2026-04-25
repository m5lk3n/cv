# Kanban Board

## Backlog

- [ ] favicon.png should be the resume.json's basic.image (use app favicon as fallback)
- [ ] commands
  - [ ] quotes
    - [x] "Whatever tomorrow brings, I'll be there. With open arms and open eyes" (Incubus, Drive)
    - [x] "The only winning move is not to play." (WarGames)
    - [ ] "Remember, you'll die." (unknown)
    - [ ] "The world doesn't need more opinions." (unknown)

## Open

- [ ] add configuration?
  - [ ] background color
  - [ ] foreground color
  - [ ] site link/title
- [ ] validate `resume/resume.json` against https://github.com/jsonresume/resume-schema/tree/master

## To do

- [ ] prep repo https://github.com/m5lk3n/cv
- [ ] https://cv.michael-klein.info/
- [ ] entry page to choose between static and dynamic
- [ ] fix package name
- [ ] rename `x-mk` to `x-cv`
- [ ] fix links
- [ ] improve wording

## Doing

- [ ] publish `resume.json` under https://registry.jsonresume.org/
  - [x] using `Makefile`
  - [ ] using GitHub workflow (see https://github.com/thomasdavis/resume)
- [ ] download resume as PDF
- [ ] add `deploy-to-web` target to Makefile
- [ ] complete `README.md`

### Done

- [x] add favicon
- [x] add link to GitHub repo
- [x] support clickable URLs
- [ ] commands
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
- [x] add license
