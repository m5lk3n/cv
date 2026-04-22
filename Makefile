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

## clean: remove generated binaries
.PHONY: clean
clean:
	rm -f mk
	rm -f web/mk.wasm

## build-cli: build the command-line interface
.PHONY: build-cli
build-cli:
	go build -o mk

## build-wasm: build the WebAssembly module
.PHONY: build-wasm
build-wasm:
	GOOS=js GOARCH=wasm go build -o web/mk.wasm

## run-localhost: build, deploy, and run the WebAssembly module on a local web server
.PHONY: run-localhost
run-localhost: needs-python3 build-wasm
	python3 -m http.server -d web