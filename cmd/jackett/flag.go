package jackett

import (
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "jackett",
	Short: "Jackett commands",
}

func init() {
	Cmd.Flags().StringP("url", "u", "http://localhost:9117", "Sonarr URL")
	Cmd.Flags().StringP("api", "a", "", "API Key")
}
