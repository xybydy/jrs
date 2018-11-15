package radarr

import (
	"github.com/spf13/cobra"
	cmd2 "jrs/cmd"
	"jrs/config"
	"jrs/pkg/radarr"
)

var (
	url string
	api string
	app *radarr.Radarr
)

var Cmd = &cobra.Command{
	Use:   "radarr",
	Short: "Radarr commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd2.CheckConfig("radarr", url, api)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		createApp()
	},
}

func createApp() {
	if app == nil {
		app = radarr.New(config.Params)
	}
}

func init() {
	url := Cmd.PersistentFlags().StringP("url", "u", "http://localhost:7878", "Radarr URL")
	api := Cmd.PersistentFlags().StringP("api", "a", "", "API Key")

	//
	// if *url != "" {
	// 	matched := utils.SplitUrl(*url)
	//
	// }
}
