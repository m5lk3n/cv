package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number and resume's last modified date",
	Run: func(cmd *cobra.Command, args []string) {
		version := "v0.1.0" // FIXME: make this dynamic based on the actual version
		r, err := resume.Load()
		if err == nil && r.Meta.LastModified != "" {
			if t, err := time.Parse(time.RFC3339, r.Meta.LastModified); err == nil {
				version += "\nresume last modified: " + t.Format("January 2, 2006")
			}
		}
		fmt.Fprintln(cmd.OutOrStdout(), version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
