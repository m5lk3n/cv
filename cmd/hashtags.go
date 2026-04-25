package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"lttl.dev/cv/resume"
)

var hashtagsCmd = &cobra.Command{
	Use:     "hashtags",
	Aliases: []string{"hashtag", "tags"},
	Short:   "Display hashtags",
	RunE: func(cmd *cobra.Command, args []string) error {
		r, err := resume.Load()
		if err != nil {
			return err
		}
		fmt.Fprintln(cmd.OutOrStdout(), strings.Join(r.XCV.Hashtags, " "))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(hashtagsCmd)
}
