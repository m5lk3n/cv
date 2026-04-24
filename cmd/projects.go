package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var projectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"proj"},
	Short:   "Display projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, p := range r.Projects {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s\n", p.Name)
			fmt.Fprintf(w, "  %s to %s\n", formatDate(p.StartDate), formatDate(p.EndDate))
			if p.URL != "" {
				fmt.Fprintf(w, "  %s\n", p.URL)
			}
			if p.Description != "" {
				fmt.Fprintf(w, "  %s\n", p.Description)
			}
			if len(p.Highlights) > 0 {
				fmt.Fprintln(w, "  Highlights:")
				for _, h := range p.Highlights {
					fmt.Fprintf(w, "    • %s\n", h)
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
