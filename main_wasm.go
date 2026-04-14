//go:build js && wasm

package main

import (
	"syscall/js"

	"lttl.dev/mk/cmd"
)

func main() {
	js.Global().Set("mkRun", js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) == 0 {
			return ""
		}
		input := args[0].String()
		return cmd.RunCommand(input)
	}))

	// Keep the Go program alive.
	select {}
}
