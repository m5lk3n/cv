package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"lttl.dev/cv/buildinfo"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number and resume's last modified date",
	Run: func(cmd *cobra.Command, args []string) {
		buildTime := buildinfo.BuildTime
		if t, err := time.Parse(time.RFC3339, buildTime); err == nil {
			buildTime = t.Format("January 2, 2006 at 15:04 MST")
		}
		version := fmt.Sprintf("CV app v%s (%s) built on %s by %s", buildinfo.Version, buildinfo.Commit, buildTime, buildinfo.BuiltBy)
		r, err := loadResume()
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
