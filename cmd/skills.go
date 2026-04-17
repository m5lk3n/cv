package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var skillsCmd = &cobra.Command{
	Use:   "skills",
	Short: "Display skills",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, s := range r.Skills {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s", s.Name)
			if s.Level != "" {
				fmt.Fprintf(w, " (%s)", s.Level)
			}
			fmt.Fprintln(w)
			if len(s.Keywords) > 0 {
				fmt.Fprintf(w, "  %s\n", strings.Join(s.Keywords, ", "))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(skillsCmd)
}
