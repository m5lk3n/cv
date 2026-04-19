package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintln(cmd.OutOrStdout(), "v0.1.0") // FIXME: make this dynamic based on the actual version
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
