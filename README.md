# mk

A CLI tool built with Go and [Cobra](https://github.com/spf13/cobra), also runnable in the browser via WebAssembly.

## Build

```sh
make build-cli       # native binary → ./mk
make build-wasm      # WebAssembly  → web/mk.wasm
```

## Run

### CLI

```sh
./mk version
```

### Web

```sh
make run-web
```

Then open http://localhost:8000 and type `mk version` in the terminal prompt.

## Clean

```sh
make clean
```

## Disclaimer

I prompted Claude Opus 4.6, 4.7, and Claude Sonnet 4.6 to implement this project. However, I had learned and used almost all technologies used in this project before, without the help of AI.