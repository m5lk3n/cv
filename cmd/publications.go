package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var publicationsCmd = &cobra.Command{
	Use:     "publications",
	Aliases: []string{"publication", "pubs", "pub", "articles", "article"},
	Short:   "Display publications",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := loadResume()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, p := range r.Publications {
			if i > 0 {
				_, _ = fmt.Fprintln(w)
			}
			_, _ = fmt.Fprintf(w, "%s\n", p.Name)
			if p.Publisher != "" {
				_, _ = fmt.Fprintf(w, "  Publisher: %s\n", p.Publisher)
			}
			if p.ReleaseDate != "" {
				_, _ = fmt.Fprintf(w, "  Date: %s\n", formatReleaseDate(p.ReleaseDate))
			}
			if p.URL != "" {
				_, _ = fmt.Fprintf(w, "  %s\n", p.URL)
			}
			if p.Summary != "" {
				_, _ = fmt.Fprintf(w, "  %s\n", p.Summary)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(publicationsCmd)
}
