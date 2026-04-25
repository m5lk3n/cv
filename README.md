# <img src="web/favicon.png" alt="Icon" height="24"> My CV & Tool

![Go](https://img.shields.io/badge/Go-1.26+-00ADDD8?logo=go&logoColor=white)
![Tool License](https://img.shields.io/badge/License-MIT-green) (App)
![!CV License](https://img.shields.io/badge/License-All%20Rights%20Reserved-lightgrey) (CV)

My CV, served as a <a href="https://github.com/spf13/cobra">Cobra</a> CLI in Go. The same binary compiles to WebAssembly so the CV runs in a browser terminal.

<a href="https://cv.michael-klein.info/"><strong>Browse CV »</strong></a>

## ✨ Features

- [x] **JSON Resume**: Follows [this](https://jsonresume.org/) open-source initiative for a JSON-based standard for resumes.
- [x] **WebAssembly**: The same Go source compiles to a native CLI and a browser-side WASM module — one command tree, two targets.
- [x] **Single binary, no runtime I/O**: CV data is embedded at compile time via `//go:embed`, so the binary (and the WASM artifact) are self-contained.
- [x] **Cobra CLI**: One subcommand per CV section ([basics](cmd/basics.go), [work](cmd/work.go), [education](cmd/education.go), [skills](cmd/skills.go), [projects](cmd/projects.go), [certificates](cmd/certificates.go), [languages](cmd/languages.go), [publications](cmd/publications.go), [references](cmd/references.go), [volunteering](cmd/volunteering.go), …).
- [x] **Browser terminal**: Keystroke-driven shell in [web/index.html](web/index.html) with command history (↑/↓), `clear`, and `reset`; URLs and email addresses are rendered as clickable links.
- [x] **MIT-licensed tooling, All-Rights-Reserved data**: Fork the tool freely; bring your own [resume.json](resume/resume.json).

### Limitation

Only English output is supported.

## 📸 Screenshots

:construction_worker:

## 🛠 Built With/On

| Area | Technology |
| :--- | :--- |
| CLI  | [Go](https://go.dev/) |
| Web  | [WebAssembly](https://webassembly.org/) |
| Data | [JSON Resume](https://jsonresume.org/) |

## 🏗 Architecture

One Go module, two entry points, one shared command tree:

- [main.go](main.go) builds a native CLI; [main_wasm.go](main_wasm.go) builds a WebAssembly module that exposes `mkRun(cmd)` to JavaScript. Build tags (`!js || !wasm` vs. `js && wasm`) pick the right entry per target.
- [cmd/](cmd/) holds the [Cobra](https://github.com/spf13/cobra) command tree — one file per CV section ([basics.go](cmd/basics.go), [work.go](cmd/work.go), [education.go](cmd/education.go), …). [root.go](cmd/root.go) exposes `Execute()` for the CLI and `RunCommand(input)` for the WASM build, which captures stdout/stderr into a buffer so the browser terminal can render it.
- [resume/](resume/) defines the [JSON Resume](https://jsonresume.org/) schema as Go structs and embeds [resume.json](resume/resume.json) at compile time via `//go:embed`. The data ships inside the binary — no filesystem reads at runtime, which is what makes the WASM build self-contained.
- [web/](web/) is the browser shell: [index.html](web/index.html) renders a terminal UI and wires keystrokes to `mkRun`; [wasm_exec.js](web/wasm_exec.js) is Go's standard WASM runtime loader; `mk.wasm` is the compiled artifact.

```text
.
├── main.go           # CLI entry (build: !js || !wasm)
├── main_wasm.go      # WASM entry (build: js && wasm)
├── cmd/              # Cobra commands — one per CV section
│   ├── root.go       # Execute() for CLI, RunCommand() for WASM
│   └── *.go          # basics, work, education, skills, …
├── resume/
│   ├── resume.go     # JSON Resume structs + //go:embed loader
│   └── resume.json   # CV content (All Rights Reserved)
└── web/              # Browser terminal hosting the WASM build
    ├── index.html
    ├── mk.wasm       # Generated using Makefile
    ├── report.pdf    # Generated using Makefile
    └── wasm_exec.js
```

## 🚀 Getting Started

To get a local copy up and running, follow these simple steps.

### Prerequisites

- Go 1.26+
- Optional: Python 3 (for the local web server)
- Optional: `jq` (for `make tidy`)
- Optional: [`gh`](https://cli.github.com/) (for `make publish-to-jsonresume`, authenticated via `gh auth login`)

`make publish-to-jsonresume` publishes [resume/resume.json](resume/resume.json) (with the `x-mk` section stripped) to a GitHub gist, making the CV available under `https://registry.jsonresume.org/<your-github-user>`. The gist ID is read from a `.env` file at the repo root:

```sh
GIST_ID=<your-gist-id>
```

`.env` is gitignored — create your own before running `make deploy`.

### Installation

```sh
git clone https://github.com/m5lk3n/cv.git
```

### Build

```sh
make build-cli       # native binary  → ./mk
make build-wasm      # WebAssembly   → web/mk.wasm
```

### Run

#### CLI

```sh
./mk                 # list available subcommands
./mk basics
./mk work
./mk education
# etc.
```

#### Web

```sh
make run-localhost
```

Open http://localhost:8000 and type commands like `mk basics` at the prompt.

### Develop

```sh
make tidy            # gofmt, go mod tidy, format resume.json
make clean           # remove ./mk and web/mk.wasm
make help            # list all targets
```

The CV content lives in [resume/resume.json](resume/resume.json) and is rendered by the subcommands in [cmd/](cmd/).

## ※ Disclaimer

I prompted Claude Opus 4.6, 4.7, and Claude Sonnet 4.6 to implement this project. However, I had learned and used almost all technologies used in this project before, without the help of AI.

## 📄 License

The code and tooling in this repository are released under the [MIT License](LICENSE).

The resume content in [resume/resume.json](resume/resume.json) is **not** covered by the MIT license — it is personal data and is **All Rights Reserved**. Please don't reuse, republish, or redistribute its contents as your own CV. Fork the tooling freely, but bring your own data.
