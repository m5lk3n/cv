package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var aboutCmd = &cobra.Command{
	Use:     "about",
	Aliases: []string{"summary"},
	Short:   "Display about summary",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		fmt.Fprintln(w, r.XMK.About)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
