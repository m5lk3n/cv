# FAQ

Q) If I deployed the web folder to a web server, would the `resume.json` in the wasm become available in search engines?

A) No. The `resume.json` is embedded in the WASM binary via Go's `//go:embed` directive. Search engine crawlers don't execute WebAssembly, so the content inside `cv.wasm` would not be visible to them. Only the static HTML in `index.html` would be indexed — and that contains no resume content, just the terminal UI shell.