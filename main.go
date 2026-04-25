//go:build !js || !wasm

package main

import "lttl.dev/cv/cmd"

func main() {
	cmd.Execute()
}
