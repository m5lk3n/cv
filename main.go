//go:build !js || !wasm

package main

import "lttl.dev/mk/cmd"

func main() {
	cmd.Execute()
}
