package cmd

import (
	"bytes"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mk",
	Short: "mk cv tool v0.1.0", // FIXME: make this dynamic based on the actual version
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

// RunCommand executes a command string and returns the captured output.
func RunCommand(input string) string {
	args := strings.Fields(input)
	if len(args) == 0 {
		return ""
	}
	// Strip leading "mk" if the user typed it.
	if args[0] == "mk" {
		args = args[1:]
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	rootCmd.Execute()
	return buf.String()
}
