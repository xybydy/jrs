package sonarr

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "sonarr",
	Short: "Sonarr commands",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	Cmd.Flags().StringP("url", "u", "http://localhost:8989", "Sonarr URL")
	Cmd.Flags().StringP("api", "a", "", "API Key")
}
