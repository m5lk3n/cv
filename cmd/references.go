package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var referencesCmd = &cobra.Command{
	Use:     "references",
	Aliases: []string{"reference", "recommendations", "recommendation", "refs", "ref"},
	Short:   "Display references/recommendations",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, ref := range r.References {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s\n", ref.Name)
			if ref.Reference != "" {
				fmt.Fprintf(w, "  %s\n", ref.Reference)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(referencesCmd)
}
