package radarr

import (
	"jrs/cmd"
	"jrs/config"
	"jrs/pkg/jackett"
	"jrs/pkg/radarr"

	"github.com/spf13/cobra"
)

var (
	url string
	api string
	app *radarr.Client
	j   *jackett.Jackett
)

var Cmd = &cobra.Command{
	Use:   "radarr",
	Short: "Radarr commands",
	Run: func(c *cobra.Command, args []string) {
		cmd.CheckConfig("radarr", url, api)
	},
	PersistentPreRun: func(c *cobra.Command, args []string) {
		if url != "" {
			config.Params.ChangeParams("Radarr", "Path", url)
		}
		if api != "" {
			config.Params.ChangeParams("Radarr", "Api", api)
		}

		createApp()
	},
}

var testIndexers = &cobra.Command{
	Use:   "test",
	Short: "Test all indexers added to Radarr",
	Run: func(c *cobra.Command, args []string) {
		app.TestAllIndexers()
	},
}

var addAllIndexers = &cobra.Command{
	Use:   "add",
	Short: "Add all available indexers to Radarr",
	Run: func(c *cobra.Command, args []string) {
		app.AddAllIndexers(j)
	},
	PreRun: func(c *cobra.Command, args []string) {
		if j == nil {
			j = jackett.New(config.Params)
		}
	},
}

var deleteAllIndexers = &cobra.Command{
	Use:   "delete",
	Short: "Delete all indexers in Radarr",
	Run: func(c *cobra.Command, args []string) {
		app.DeleteAllIndexers()
	},
}

func createApp() {
	if app == nil {
		app = radarr.NewClient()
	}
}

func init() {
	Cmd.PersistentFlags().StringVarP(&url, "url", "u", "", "Radarr URL")
	Cmd.PersistentFlags().StringVarP(&api, "api", "a", "", "API Key")

	Cmd.AddCommand(testIndexers, addAllIndexers, deleteAllIndexers)

}
