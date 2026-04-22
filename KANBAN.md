# Kanban Board

## Open

- [ ] add configuration?
  - [ ] background color
  - [ ] foreground color
  - [ ] site link/title
- [ ] validate `resume/resume.json` against https://github.com/jsonresume/resume-schema/tree/master

## To do

- [ ] favicon.png should be the resume.json's basic.image (use app favicon as fallback)
- [ ] prep repo https://github.com/m5lk3n/cv
- [ ] download resume as PDF
- [ ] entry page to choose between static and dynamic
- [ ] add deploy target to Makefile
- [ ] add license
- [ ] update `README.md`
- [ ] publish on JSON resume's registry (see also https://github.com/thomasdavis/resume)
- [ ] fix links
- [ ] fix package name
- [ ] address FIXMEs
- [ ] improve wording

## Doing

### Commands

- [ ] about / summary
- [ ] experience / work
  - [ ] achievements
- [ ] skills
- [ ] certification(s) (license(s) as alias?)
- [ ] projects
- [ ] basics
  - [ ] pages / links / online / social
- [ ] version
  - [ ] embed semver in binary, read version from there
- [ ] education
- [ ] language(s)
- [ ] volunteering
- [ ] quotes
  - [x] "Whatever tomorrow brings, I'll be there. With open arms and open eyes" (Incubus, Drive)
  - [x] "The only winning move is not to play." (WarGames)
  - [ ] "Remember, you'll die." (unknown)
  - [ ] "The world doesn't need more opinions." (unknown)
- [ ] faq
  - [ ] recruiter
  - [ ] sales

### Done

- [x] add favicon
- [x] add link to GitHub repo
- [x] support clickable URLs
- [x] register `achievement(s)` as top-level command for `work achievement(s)`