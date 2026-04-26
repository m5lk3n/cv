//go:build js && wasm

package main

import (
	"syscall/js"

	"lttl.dev/cv/buildinfo"
	"lttl.dev/cv/cmd"
)

func main() {
	js.Global().Set("cvRun", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 {
			return ""
		}
		input := args[0].String()
		return cmd.RunCommand(input)
	}))

	js.Global().Set("cvName", js.FuncOf(func(this js.Value, args []js.Value) any {
		return buildinfo.NameOnCV
	}))

	js.Global().Set("ghUsername", js.FuncOf(func(this js.Value, args []js.Value) any {
		return buildinfo.GitHubUsername
	}))

	// Keep the Go program alive.
	select {}
}
