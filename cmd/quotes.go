package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var quotesCmd = &cobra.Command{
	Use:     "quotes",
	Aliases: []string{"quote"},
	Short:   "Display quotes",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for _, q := range r.XCV.Quotes {
			fmt.Fprintf(w, "\"%s\"\n", q.Text)
			if q.Author != "" {
				fmt.Fprintf(w, "  — %s\n", q.Author)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(quotesCmd)
}
