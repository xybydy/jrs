package sonarr

import (
	"github.com/xybydy/jrs/cmd"
	"github.com/xybydy/jrs/config"
	"github.com/xybydy/jrs/pkg/jackett"
	"github.com/xybydy/jrs/pkg/trackers"
	"github.com/xybydy/jrs/pkg/trackers/sonarr"

	"github.com/spf13/cobra"
)

var (
	url string
	api string
	app *trackers.Client
	j   *jackett.Jackett
)

var Cmd = &cobra.Command{
	Use:   "sonarr",
	Short: "Sonarr commands",
	Run: func(c *cobra.Command, args []string) {
		cmd.CheckConfig("sonarr", url, api)
	},
	PersistentPreRun: func(c *cobra.Command, args []string) {
		if url != "" {
			config.Params.ChangeParams("Sonarr", "path", url)
		}
		if api != "" {
			config.Params.ChangeParams("Sonarr", "api", api)
		}

		createApp()
	},
}

var testIndexers = &cobra.Command{
	Use:   "test",
	Short: "Test all indexers added",
	Run: func(c *cobra.Command, args []string) {
		app.TestAllIndexers()
	},
}

var addAllIndexers = &cobra.Command{
	Use:   "add",
	Short: "Add all available indexers",
	Run: func(c *cobra.Command, args []string) {
		if len(args) != 0 {
			app.AddAllIndexers(j, args[0])
		} else {
			app.AddAllIndexers(j, "")
		}
	},
	PreRun: func(c *cobra.Command, args []string) {
		if j == nil {
			j = jackett.New(config.Params)
		}
	},
}

var deleteAllIndexers = &cobra.Command{
	Use:   "delete",
	Short: "Delete all indexers",
	Run: func(c *cobra.Command, args []string) {
		app.DeleteAllIndexers()
	},
}

func createApp() {
	if app == nil {
		app = sonarr.NewClient()
	}
}

func init() {
	Cmd.Flags().StringVarP(&url, "url", "u", "", "Sonarr URL")
	Cmd.Flags().StringVarP(&api, "api", "a", "", "API Key")

	Cmd.AddCommand(testIndexers, addAllIndexers, deleteAllIndexers)
}
