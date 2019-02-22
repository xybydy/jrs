// CLI commands for configuration
package config

import (
	"jrs/config"
	"log"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var (
	url       string
	api       string
	flagExist bool
)

// Command of configuration
var Config = &cobra.Command{
	Use:   "config",
	Short: "Command of the config parameters",
}

var set = &cobra.Command{
	Use:     "set",
	Short:   "Set of commands to save the config parameters to the config file formatted as ´toml´ for each application.",
	Example: "jrs config set [command] [flags] [filename]",
}

func saveSettings(app string, args []string) {
	configFile := config.DEFAULT_CONFIG_PATH

	if url == "" && api == "" {
		log.Fatalf("There is no config parameters passed. Nothing to save!")
	}

	if url != "" {
		config.Params.ChangeParams(strings.Title(app), "Path", url)
	}
	if api != "" {
		config.Params.ChangeParams(strings.Title(app), "Api", api)
	}

	if len(args) > 0 {
		configFile = args[0]
		log.Printf("Saving config file as %s.", configFile)
	} else {
		log.Printf("There is no config file specified. Saving config file as %s.", configFile)
	}

	err := config.Params.SaveFile(configFile)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func isFlagPassed(cmd *cobra.Command) {
	checker := func(f *pflag.Flag) {
		if f.Changed {
			flagExist = true
		}
	}
	cmd.LocalFlags().VisitAll(checker)
}

var radarr = &cobra.Command{
	Use:     "radarr",
	Short:   "Radarr command to pass and the save the parameters to config file.",
	Example: "jrs config set radarr --api jashdakj1827jfhs --url http://localhost:8888 settings.toml",
	Run: func(cmd *cobra.Command, args []string) {

		if isFlagPassed(cmd); !flagExist {
			if err := cmd.Help(); err != nil {
				log.Fatalln(err)
			}
		} else {
			saveSettings("radarr", args)
		}
	},
}

var sonarr = &cobra.Command{
	Use:   "sonarr",
	Short: "Sonarr command to pass and the save the parameters to config file.",
	Run: func(cmd *cobra.Command, args []string) {

		if isFlagPassed(cmd); !flagExist {
			if err := cmd.Help(); err != nil {
				log.Fatalln(err)
			}
		} else {
			saveSettings("sonarr", args)
		}
	},
}

var jackett = &cobra.Command{
	Use:   "jackett",
	Short: "Jackett command to pass and the save the parameters to config file.",
	Run: func(cmd *cobra.Command, args []string) {

		if isFlagPassed(cmd); !flagExist {
			if err := cmd.Help(); err != nil {
				log.Fatalln(err)
			}
		} else {
			saveSettings("jackett", args)
		}
	},
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
