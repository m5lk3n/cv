package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var languagesCmd = &cobra.Command{
	Use:     "languages",
	Aliases: []string{"language", "langs", "lang"},
	Short:   "Display languages",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := loadResume()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for _, lang := range r.Languages {
			fmt.Fprintf(w, "%s — %s\n", lang.Language, lang.Fluency)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(languagesCmd)
}
