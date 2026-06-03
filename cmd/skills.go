package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var skillsCmd = &cobra.Command{
	Use:     "skills",
	Aliases: []string{"skill"},
	Short:   "Display skills",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := loadResume()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, s := range r.Skills {
			if i > 0 {
				_, _ = fmt.Fprintln(w)
			}
			_, _ = fmt.Fprintf(w, "%s", s.Name)
			if s.Level != "" {
				_, _ = fmt.Fprintf(w, " (%s)", s.Level)
			}
			_, _ = fmt.Fprintln(w)
			if len(s.Keywords) > 0 {
				_, _ = fmt.Fprintf(w, "  %s\n", strings.Join(s.Keywords, ", "))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(skillsCmd)
}
