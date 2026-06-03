package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var certificatesCmd = &cobra.Command{
	Use:     "certificates",
	Aliases: []string{"certifications", "certs", "cert", "certificate", "certification", "licenses", "license", "lic"},
	Short:   "Display certificates",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := loadResume()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, c := range r.Certificates {
			if i > 0 {
				_, _ = fmt.Fprintln(w)
			}
			_, _ = fmt.Fprintf(w, "%s\n", c.Name)
			if c.Issuer != "" {
				_, _ = fmt.Fprintf(w, "  Issuer: %s\n", c.Issuer)
			}
			if c.Date != "" {
				_, _ = fmt.Fprintf(w, "  Date: %s\n", formatDate(c.Date))
			}
			if c.URL != "" {
				_, _ = fmt.Fprintf(w, "  %s\n", c.URL)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(certificatesCmd)
}
