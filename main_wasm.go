//go:build js && wasm

package main

import (
	"syscall/js"

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

	// Keep the Go program alive.
	select {}
}
