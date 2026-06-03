package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var projectsCmd = &cobra.Command{
	Use:     "projects",
	Aliases: []string{"proj"},
	Short:   "Display projects",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := loadResume()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, p := range r.Projects {
			if i > 0 {
				_, _ = fmt.Fprintln(w)
			}
			_, _ = fmt.Fprintf(w, "%s\n", p.Name)
			_, _ = fmt.Fprintf(w, "  %s to %s\n", formatDate(p.StartDate), formatEndDate(p.EndDate))
			if p.URL != "" {
				_, _ = fmt.Fprintf(w, "  %s\n", p.URL)
			}
			if p.Description != "" {
				_, _ = fmt.Fprintf(w, "  %s\n", p.Description)
			}
			if len(p.Highlights) > 0 {
				_, _ = fmt.Fprintln(w, "  Highlights:")
				for _, h := range p.Highlights {
					_, _ = fmt.Fprintf(w, "    • %s\n", h)
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(projectsCmd)
}
