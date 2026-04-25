package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"lttl.dev/cv/resume"
)

var certificatesCmd = &cobra.Command{
	Use:     "certificates",
	Aliases: []string{"certifications", "certs", "cert", "certificate", "certification", "licenses", "license", "lic"},
	Short:   "Display certificates",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		w := cmd.OutOrStdout()
		for i, c := range r.Certificates {
			if i > 0 {
				fmt.Fprintln(w)
			}
			fmt.Fprintf(w, "%s\n", c.Name)
			if c.Issuer != "" {
				fmt.Fprintf(w, "  Issuer: %s\n", c.Issuer)
			}
			if c.Date != "" {
				fmt.Fprintf(w, "  Date: %s\n", formatDate(c.Date))
			}
			if c.URL != "" {
				fmt.Fprintf(w, "  %s\n", c.URL)
			}
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(certificatesCmd)
}
