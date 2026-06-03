package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var basicsCmd = &cobra.Command{
	Use:     "basics",
	Aliases: []string{"contact"},
	Short:   "Display the basics including contact info",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := loadResume()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		b := r.Basics
		_, _ = fmt.Fprintf(w, "%s — %s\n", b.Name, b.Label)
		if b.Image != "" {
			_, _ = fmt.Fprintf(w, "  Image: %s\n", b.Image)
		}
		if b.Email != "" {
			_, _ = fmt.Fprintf(w, "  Email: %s\n", b.Email)
		}
		if b.Phone != "" {
			_, _ = fmt.Fprintf(w, "  Phone: %s\n", b.Phone)
		}
		if b.URL != "" {
			_, _ = fmt.Fprintf(w, "  Web: %s\n", b.URL)
		}
		if b.Summary != "" {
			_, _ = fmt.Fprintf(w, "  Summary: %s\n", b.Summary)
		}
		loc := b.Location
		var parts []string
		if loc.City != "" {
			parts = append(parts, loc.City)
		}
		if loc.Region != "" {
			parts = append(parts, loc.Region)
		}
		if loc.CountryCode != "" {
			parts = append(parts, loc.CountryCode)
		}
		if len(parts) > 0 {
			_, _ = fmt.Fprintf(w, "  Location: %s\n", strings.Join(parts, ", "))
		}
		for _, p := range b.Profiles {
			_, _ = fmt.Fprintf(w, "  %s: %s\n", p.Network, p.URL)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(basicsCmd)
}
