package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var aboutCmd = &cobra.Command{
	Use:     "about",
	Aliases: []string{"summary", "tldr", "info"},
	Short:   "Display about summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := loadResume()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		fmt.Fprintln(w, r.Basics.Summary)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
