package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var faqsCmd = &cobra.Command{
	Use:     "faqs",
	Aliases: []string{"faq"},
	Short:   "Display FAQs",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, f := range r.XMK.FAQs {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "Q: %s\n", f.Question)
			fmt.Fprintf(w, "A: %s\n", f.Answer)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(faqsCmd)
}
