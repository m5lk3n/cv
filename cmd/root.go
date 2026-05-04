package cmd

import (
	"bytes"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"lttl.dev/cv/buildinfo"
)

var rootCmd = &cobra.Command{
	Use:   "cv",
	Short: buildinfo.NameOnCV + "'s CV",
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
	// Strip leading "cv" if the user typed it.
	if args[0] == "cv" {
		args = args[1:]
	}

	buf := new(bytes.Buffer)
	rootCmd.SetOut(buf)
	rootCmd.SetErr(buf)
	rootCmd.SetArgs(args)
	if err := rootCmd.Execute(); err != nil {
		return err.Error()
	}
	return buf.String()
}
