ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ifeq (,$(wildcard ./.env))
$(error .env file not found. Please create .env (or copy .env.example))
endif

VERSION    := $(shell git describe --tags --abbrev=0)
COMMIT     := $(shell git rev-parse --short HEAD)
BUILD_TIME := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
BUILT_BY   := $(shell echo $$USER)

LDFLAGS := -X 'lttl.dev/mk/buildinfo.Version=$(VERSION)' \
           -X 'lttl.dev/mk/buildinfo.Commit=$(COMMIT)' \
           -X 'lttl.dev/mk/buildinfo.BuildTime=$(BUILD_TIME)' \
           -X 'lttl.dev/mk/buildinfo.BuiltBy=$(BUILT_BY)'

## help: print this help message
.PHONY: help
help:
	@echo 'usage: make <target>'
	@echo
	@echo '  where <target> is one of the following:'
	@echo
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

needs-python3:
	@command -v python3 >/dev/null 2>&1 || { echo >&2 "python3 is required but it's not installed. Aborting."; exit 1; }

needs-jq:
	@command -v jq >/dev/null 2>&1 || { echo >&2 "jq is required but it's not installed. Aborting."; exit 1; }

needs-gh:
	@command -v gh >/dev/null 2>&1 || { echo >&2 "gh is required but it's not installed. Aborting."; exit 1; }

needs-chromium:
	@command -v chromium >/dev/null 2>&1 || { echo >&2 "chromium is required but it's not installed. Aborting."; exit 1; }

## clean: remove generated files
.PHONY: clean
clean:
	rm -rf dist
	rm -f mk
	rm -f web/mk.wasm
	rm -f web/resume.pdf

## tidy: format Go files, tidy Go modules, and format resume.json
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

## build-web: build the WebAssembly module and export the PDF for web usage
.PHONY: build-web
build-web: build-wasm export-pdf
	@echo "-=> WebAssembly module built and PDF exported to web/resume.pdf"

## run-localhost: build, deploy, and run the WebAssembly module on a local web server
.PHONY: run-localhost
run-localhost: needs-python3 build-web
	@python3 -m http.server -d web

## publish-to-jsonresume: make the resume available under registry.jsonresume.org
.PHONY: publish-to-jsonresume
publish-to-jsonresume: needs-jq needs-gh
	@mkdir -p dist
	@jq 'del(."x-cv")' resume/resume.json > dist/resume.json
	@gh gist edit $(GIST_ID) -f resume.json dist/resume.json
	@GITHUB_USERNAME=$$(gh api user --jq .login); \
	echo "Updated resume available under https://registry.jsonresume.org/$$GITHUB_USERNAME"

## export-pdf: export the resume as a PDF using headless Chromium
.PHONY: export-pdf
export-pdf: needs-gh needs-chromium
	@GITHUB_USERNAME=$$(gh api user --jq .login); \
	chromium --headless --no-pdf-header-footer --print-to-pdf=web/resume.pdf https://registry.jsonresume.org/$$GITHUB_USERNAME

## publish-to-web: build and make the web directory available online
.PHONY: publish-to-web
publish-to-web: build-web
	rsync -vz --delete --recursive web/ "${DEPLOY_TARGET}"
	@echo "-=> Check file permissions under ${DEPLOY_TARGET}"