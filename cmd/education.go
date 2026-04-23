package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var educationCmd = &cobra.Command{
	Use:     "education",
	Aliases: []string{"edu"},
	Short:   "Display education history",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, edu := range r.Education {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s — %s (%s)\n", edu.Institution, edu.Area, edu.StudyType)
			fmt.Fprintf(w, "  %s to %s\n", edu.StartDate, edu.EndDate)
			if edu.URL != "" {
				fmt.Fprintf(w, "  %s\n", edu.URL)
			}
			if edu.Score != "" {
				fmt.Fprintf(w, "  GPA: %s\n", edu.Score)
			}
			if len(edu.Courses) > 0 {
				fmt.Fprintf(w, "  Courses: %s\n", strings.Join(edu.Courses, ", "))
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(educationCmd)
}
