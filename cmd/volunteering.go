package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var volunteeringCmd = &cobra.Command{
	Use:     "volunteering",
	Aliases: []string{"volunteer", "vol"},
	Short:   "Display volunteering experience",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, v := range r.Volunteer {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s — %s\n", v.Organization, v.Position)
			fmt.Fprintf(w, "  %s to %s\n", v.StartDate, v.EndDate)
			if v.URL != "" {
				fmt.Fprintf(w, "  %s\n", v.URL)
			}
			if v.Summary != "" {
				fmt.Fprintf(w, "  %s\n", v.Summary)
			}
			if len(v.Highlights) > 0 {
				fmt.Fprintln(w, "  Highlights:")
				for _, h := range v.Highlights {
					fmt.Fprintf(w, "    • %s\n", h)
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(volunteeringCmd)
}
