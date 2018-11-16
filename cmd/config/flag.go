package config

import (
	"jrs/config"

	"github.com/spf13/cobra"
)

var (
	url string
	api string
)

var Config = &cobra.Command{
	Use:   "config",
	Short: "Configuration",
}

var set = &cobra.Command{
	Use: "set",
}

var radarr = &cobra.Command{
	Use: "radarr",
	Run: func(cmd *cobra.Command, args []string) {
		if url != "" {
			config.Params.ChangeParams("radarr", "path", url)
		}
		if api != "" {
			config.Params.ChangeParams("radarr", "api", api)
		}
		config.Params.SaveFile(args[0])

	},
	Args: cobra.ExactArgs(1),
}

var sonarr = &cobra.Command{
	Use: "sonarr",
	Run: func(cmd *cobra.Command, args []string) {
		if url != "" {
			config.Params.ChangeParams("sonarr", "path", url)
		}
		if api != "" {
			config.Params.ChangeParams("sonarr", "api", api)
		}
		config.Params.SaveFile(args[0])

	},
	Args: cobra.ExactArgs(1),
}

var jackett = &cobra.Command{
	Use: "jackett",
	Run: func(cmd *cobra.Command, args []string) {
		if url != "" {
			config.Params.ChangeParams("jackett", "path", url)
		}
		if api != "" {
			config.Params.ChangeParams("jackett", "api", api)
		}
		config.Params.SaveFile(args[0])

	},
	Args: cobra.ExactArgs(1),
}

func init() {
	radarr.Flags().StringVarP(&url, "url", "u", "",
		"Full address of the application server, i.e. http://192.168.1.1:7878")
	radarr.Flags().StringVarP(&api, "api", "a", "",
		"API Key of the service.")
	sonarr.Flags().StringVarP(&url, "url", "u", "",
		"Full address of the application server, i.e. http://192.168.1.1:8989")
	sonarr.Flags().StringVarP(&api, "api", "a", "",
		"API Key of the service.")
	jackett.Flags().StringVarP(&url, "url", "u", "",
		"Full address of the application server, i.e. http://192.168.1.1:9117")
	jackett.Flags().StringVarP(&api, "api", "a", "",
		"API Key of the service.")

	set.AddCommand(radarr, sonarr, jackett)
	Config.AddCommand(set)
}
