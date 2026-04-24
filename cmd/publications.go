package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var publicationsCmd = &cobra.Command{
	Use:     "publications",
	Aliases: []string{"publication", "pubs", "pub", "articles", "article"},
	Short:   "Display publications",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, p := range r.Publications {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s\n", p.Name)
			if p.Publisher != "" {
				fmt.Fprintf(w, "  Publisher: %s\n", p.Publisher)
			}
			if p.ReleaseDate != "" {
				fmt.Fprintf(w, "  Date: %s\n", formatReleaseDate(p.ReleaseDate))
			}
			if p.URL != "" {
				fmt.Fprintf(w, "  %s\n", p.URL)
			}
			if p.Summary != "" {
				fmt.Fprintf(w, "  %s\n", p.Summary)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(publicationsCmd)
}
