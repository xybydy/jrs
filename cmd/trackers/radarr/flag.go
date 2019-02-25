package radarr

import (
	"jrs/cmd"
	"jrs/config"
	"jrs/pkg/jackett"
	"jrs/pkg/trackers/radarr"

	"github.com/spf13/cobra"
)

var (
	url string
	api string
	app *radarr.Radarr
	j   *jackett.Jackett
)

func createApp() {
	if app == nil {
		app = radarr.New(config.Params)
	}
}

var Cmd = &cobra.Command{
	Use:   "radarr",
	Short: "Radarr commands",
	Run: func(c *cobra.Command, args []string) {
		cmd.GetConfig("radarr", url, api)
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

var TestIndexers = &cobra.Command{
	Use:   "test",
	Short: "Test all indexers added to Radarr",
	Run: func(c *cobra.Command, args []string) {
		app.TestAllIndexers()
	},
}

var AddAllIndexers = &cobra.Command{
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

var DeleteAllIndexers = &cobra.Command{
	Use:   "delete",
	Short: "Delete all indexers in Radarr",
	Run: func(c *cobra.Command, args []string) {
		app.DeleteAllIndexers()
	},
}

// var BulkImportMovies = &cobra.Command{
// 	User:  "bulkimport",
// 	Short: "Scans folder provided to import the movies",
// 	Run: func(c *cobra.Command, args []string) {
// 		if len(args < 1) {
// 			log.Fatalf("There is no folder provided")
// 		}
// 	},
// }

func init() {
	Cmd.PersistentFlags().StringVarP(&url, "url", "u", "", "Radarr URL")
	Cmd.PersistentFlags().StringVarP(&api, "api", "a", "", "API Key")

	Cmd.AddCommand(TestIndexers, AddAllIndexers, DeleteAllIndexers)

}
