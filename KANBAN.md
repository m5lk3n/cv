# Kanban Board

## Backlog

- [ ] support additional themes?
- [ ] clarify license/distribution `web/wasm_exec.js` 
- [ ] favicon.png should be the resume.json's basic.image (use app favicon as fallback)
- [ ] commands
  - [ ] quotes
    - [x] "Whatever tomorrow brings, I'll be there. With open arms and open eyes" (Incubus, Drive)
    - [x] "The only winning move is not to play." (WarGames)
    - [ ] "Remember, you'll die." (unknown)
    - [ ] "The world doesn't need more opinions." (unknown)
- [ ] check out https://github.com/igolaizola/webcli

## Open

- [ ] add configuration?
  - [ ] background color
  - [ ] foreground color
  - [ ] links
  - [ ] site link/title

## To do

- [ ] add `.env.example`
- [ ] prep repo https://github.com/m5lk3n/cv
- [ ] entry page to choose between static and dynamic
- [ ] fix package name
- [ ] fix links incl. hardcodings like https://registry.jsonresume.org/m5lk3n in `index.html`
- [ ] improve wording

## Doing

- [ ] publish `resume.json` under https://registry.jsonresume.org/
  - [x] using `Makefile`
  - [ ] using GitHub workflow (see https://github.com/thomasdavis/resume)
- [ ] complete `README.md`, also include screenshots
  - [ ] add `screenshots/animated.gif`

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
- [x] https://cv.michael-klein.info/
- [x] resume as downloadable PDF
- [x] add link to published resume under registry.jsonresume.org
- [x] add `publish-to-web` target to Makefile
