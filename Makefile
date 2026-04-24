VERSION    := $(shell git describe --tags --abbrev=0)
COMMIT     := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
BUILT_BY   := $(shell echo $$USER)

LDFLAGS := -X 'lttl.dev/mk/app.Version=$(VERSION)' \
           -X 'lttl.dev/mk/app.Commit=$(COMMIT)' \
           -X 'lttl.dev/mk/app.BuildTime=$(BUILD_TIME)' \
           -X 'lttl.dev/mk/app.BuiltBy=$(BUILT_BY)'

## help: print this help message
.PHONY: help
help:
	@echo 'usage: make <target>'
	@echo
	@echo '  where <target> is one of the following:'
	@echo
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

needs-python3:
	@command -v python3 >/dev/null 2>&1 || { echo >&2 "Python 3 is required but it's not installed. Aborting."; exit 1; }

needs-jq:
	@command -v jq >/dev/null 2>&1 || { echo >&2 "jq is required but it's not installed. Aborting."; exit 1; }

## clean: remove generated binaries
.PHONY: clean
clean:
	rm -f mk
	rm -f web/mk.wasm

## tidy: tidy up the codebase by formatting Go files, tidying Go modules, and formatting JSON in the resume file
.PHONY: tidy
tidy: needs-jq
	gofmt -s -w .
	go mod tidy
	jq . resume/resume.json > resume/resume.json.tmp && mv resume/resume.json.tmp resume/resume.json

## build-cli: build the command-line interface
.PHONY: build-cli
build-cli:
	go build -o mk -ldflags="$(LDFLAGS)"

## build-wasm: build the WebAssembly module
.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o web/mk.wasm -ldflags="$(LDFLAGS)"

## run-localhost: build, deploy, and run the WebAssembly module on a local web server
.PHONY: run-localhost
run-localhost: needs-python3 build-wasm
	python3 -m http.server -d web

## deploy: write resume.json without the x-mk section to dist/resume.json
.PHONY: deploy
deploy: needs-jq
	mkdir -p dist
	jq 'del(."x-mk")' resume/resume.json > dist/resume.json