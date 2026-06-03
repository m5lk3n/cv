package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var faqsCmd = &cobra.Command{
	Use:     "faqs",
	Aliases: []string{"faq"},
	Short:   "Display FAQs",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := loadResume()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, f := range r.XCV.FAQs {
			if i > 0 {
				_, _ = fmt.Fprintln(w)
			}
			_, _ = fmt.Fprintf(w, "Q: %s\n", f.Question)
			_, _ = fmt.Fprintf(w, "A: %s\n", f.Answer)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(faqsCmd)
}
