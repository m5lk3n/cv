package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/mk/resume"
)

var workCmd = &cobra.Command{
	Use:     "work",
	Aliases: []string{"experience", "exp"},
	Short:   "Display work experience",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, job := range r.Work {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s — %s\n", job.Name, job.Position)
			fmt.Fprintf(w, "  %s to %s\n", job.StartDate, job.EndDate)
			if job.URL != "" {
				fmt.Fprintf(w, "  %s\n", job.URL)
			}
			if job.Summary != "" {
				fmt.Fprintf(w, "  %s\n", job.Summary)
			}
			if len(job.Highlights) > 0 {
				fmt.Fprintln(w, "  Highlights:")
				for _, h := range job.Highlights {
					fmt.Fprintf(w, "    • %s\n", h)
				}
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(workCmd)
	workCmd.AddCommand(achievementsCmd)
}

var achievementsCmd = &cobra.Command{
	Use:     "achievements",
	Aliases: []string{"achieved", "achievement"},
	Short:   "List all work achievements",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, job := range r.Work {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s:\n", job.Name)
			for _, h := range job.Highlights {
				fmt.Fprintf(w, "  • %s\n", h)
			}
		}
		return nil
	},
}
