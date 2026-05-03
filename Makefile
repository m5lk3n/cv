ifneq (,$(wildcard ./.env))
    include .env
    export
endif

ifeq (,$(wildcard ./.env))
$(error .env file not found. Please create .env (or copy .env.example))
endif

VERSION          := $(shell git describe --tags --abbrev=0)
COMMIT           := $(shell git rev-parse --short HEAD)
BUILD_TIME       := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)
BUILT_BY         := $(shell echo $$USER)
GITHUB_USERNAME  := $(shell gh api user --jq .login 2>/dev/null)

LDFLAGS := -X 'lttl.dev/cv/buildinfo.Version=$(VERSION)' \
           -X 'lttl.dev/cv/buildinfo.Commit=$(COMMIT)' \
           -X 'lttl.dev/cv/buildinfo.BuildTime=$(BUILD_TIME)' \
           -X 'lttl.dev/cv/buildinfo.BuiltBy=$(BUILT_BY)' \
           -X 'lttl.dev/cv/buildinfo.NameOnCV=$(NAME_ON_CV)' \
           -X 'lttl.dev/cv/buildinfo.GitHubUsername=$(GITHUB_USERNAME)'

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

needs-swag:
	@command -v swag >/dev/null 2>&1 || { echo >&2 "swag is required but it's not installed (use `go install github.com/swaggo/swag/cmd/swag@latest` to install it). Aborting."; exit 1; }

.PHONY: assets
assets: api/favicon.png

api/favicon.png: web/favicon.png
	cp web/favicon.png api/favicon.png

## install-githooks: point git at the versioned hooks under .githooks
.PHONY: install-githooks
install-githooks:
	git config core.hooksPath .githooks
	@echo "Git hooks path set to .githooks"

## clean: remove generated files
.PHONY: clean
clean:
	rm -rf dist
	rm -rf bin
	rm -f web/cv.wasm
	rm -f web/resume.pdf

## tidy: format Go files, tidy Go modules, and format resume.json
.PHONY: tidy
tidy: needs-jq
	gofmt -s -w .
	go mod tidy
	jq . resume/resume.json > resume/resume.json.tmp && mv resume/resume.json.tmp resume/resume.json

## test: run the Go test suite
.PHONY: test
test:
	go test ./...

## build-cli: build the command-line interface
.PHONY: build-cli
build-cli:
	@mkdir -p bin
	go build -o bin/cv -ldflags="$(LDFLAGS)"

## build-api: build the HTTP API server and regenerate Swagger docs
.PHONY: build-api
build-api: assets generate-apidoc
	@mkdir -p bin
	go build -o bin/cvapi -ldflags="$(LDFLAGS)" ./cvapi

## generate-apidoc: regenerate Swagger docs
.PHONY: generate-apidoc
generate-apidoc: needs-swag
	swag init -g cvapi/main.go -d . -o docs --parseInternal

## build-wasm: build the WebAssembly module
.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o web/cv.wasm -ldflags="$(LDFLAGS)"

## build-web: build the WebAssembly module and export the PDF for web usage
.PHONY: build-web
build-web: build-wasm export-pdf
	@echo "WebAssembly module built and PDF exported to web/resume.pdf"

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
	echo "Updated CV published under https://registry.jsonresume.org/$$GITHUB_USERNAME"

## export-pdf: export the registry's resume as a PDF using headless Chromium
.PHONY: export-pdf
export-pdf: needs-gh needs-chromium
	@GITHUB_USERNAME=$$(gh api user --jq .login); \
	chromium --log-level=3 --headless --no-pdf-header-footer --print-to-pdf=web/resume.pdf https://registry.jsonresume.org/$$GITHUB_USERNAME

## publish-to-web: build and make the web directory available online
.PHONY: publish-to-web
publish-to-web: build-web
	rsync -vz --delete --recursive --chown=caddy:caddy web/ "${DEPLOY_TARGET}"
	@echo "Updated CV app version published to web"

## publish: build, export PDF, publish to JSONResume and web
.PHONY: publish
publish: tidy publish-to-jsonresume publish-to-web