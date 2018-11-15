package jackett

import (
	"github.com/spf13/cobra"
	cmd2 "jrs/cmd"
	"jrs/config"
	"jrs/pkg/jackett"
)

var (
	url  string
	api  string
	user string
	id   string
	pass string
	app  *jackett.Jackett
)

func createApp() {
	if app == nil {
		app = jackett.New(config.Params)
	}
}

var Cmd = &cobra.Command{
	Use:   "jackett",
	Short: "Jackett commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd2.CheckConfig("jackett", url, api)
	},
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		createApp()
	},
}

var getAllIndexers = &cobra.Command{
	Use:   "get",
	Short: "Command to get all configured indexers.",
	Run: func(cmd *cobra.Command, args []string) {
		app.GetConfiguredIndexers()
	},
}

var listIndexers = &cobra.Command{
	Use:   "list",
	Short: "Command to get all indexers in jackett",
	Run: func(cmd *cobra.Command, args []string) {
		app.GetAllIndexers()
	},
}

var addIndexer = &cobra.Command{
	Use:   "add",
	Short: "Command to add indexer in jackett",
	Long:  "Please use get command to get id of the indexer before adding it.",
	Run: func(cmd *cobra.Command, args []string) {
		app.AddIndexer(id, user, pass)
	},
}

var addAllPublicIndexers = &cobra.Command{
	Use:   "addall",
	Short: "Command to add all public indexers to jackett",
	Run: func(cmd *cobra.Command, args []string) {
		app.AddAllPublicIndexers()
	},
}

func init() {
	Cmd.PersistentFlags().StringVarP(&url, "url", "u", "http://localhost:9117", "Jackett URL")
	Cmd.PersistentFlags().StringVarP(&api, "api", "a", "", "API Key")

	addIndexer.Flags().StringVarP(&user, "username", "s", "", "Tracker Username")
	addIndexer.Flags().StringVarP(&id, "id", "i", "", "Tracker ID")
	addIndexer.Flags().StringVarP(&pass, "password", "p", "", "Tracker Password")

	Cmd.AddCommand(addAllPublicIndexers)
	Cmd.AddCommand(getAllIndexers)
	Cmd.AddCommand(listIndexers)
	Cmd.AddCommand(addIndexer)

}
