package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"lttl.dev/mk/buildinfo"
	"lttl.dev/mk/resume"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number and resume's last modified date",
	Run: func(cmd *cobra.Command, args []string) {
		buildTime := buildinfo.BuildTime
		if t, err := time.Parse(time.RFC3339, buildTime); err == nil {
			buildTime = t.Format("January 2, 2006 at 15:04 MST")
		}
		version := fmt.Sprintf("App built on %s by %s", buildTime, buildinfo.BuiltBy)
		version += fmt.Sprintf("\nApp v%s (%s)", buildinfo.Version, buildinfo.Commit)
		r, err := resume.Load()
		if err == nil && r.Meta.LastModified != "" {
			if t, err := time.Parse(time.RFC3339, r.Meta.LastModified); err == nil {
				version += "\nCV data last modified on " + t.Format("January 2, 2006")
			}
		}
		fmt.Fprintln(cmd.OutOrStdout(), version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
