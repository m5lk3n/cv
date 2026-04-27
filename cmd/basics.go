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
		fmt.Fprintf(w, "%s — %s\n", b.Name, b.Label)
		if b.Image != "" {
			fmt.Fprintf(w, "  Image: %s\n", b.Image)
		}
		if b.Email != "" {
			fmt.Fprintf(w, "  Email: %s\n", b.Email)
		}
		if b.Phone != "" {
			fmt.Fprintf(w, "  Phone: %s\n", b.Phone)
		}
		if b.URL != "" {
			fmt.Fprintf(w, "  Web: %s\n", b.URL)
		}
		if b.Summary != "" {
			fmt.Fprintf(w, "  Summary: %s\n", b.Summary)
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
			fmt.Fprintf(w, "  Location: %s\n", strings.Join(parts, ", "))
		}
		for _, p := range b.Profiles {
			fmt.Fprintf(w, "  %s: %s\n", p.Network, p.URL)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(basicsCmd)
}
