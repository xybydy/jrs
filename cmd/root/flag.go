package root

import (
	cfg "jrs/cmd/config"
	"jrs/cmd/jackett"
	"jrs/cmd/radarr"
	"jrs/cmd/sonarr"
	"jrs/config"
	"jrs/utils"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{Use: "jrs",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&config.ConfPath, "config", "c", config.DEFAULT_CONFIG_PATH, "Config file path")
	RootCmd.AddCommand(radarr.Cmd, sonarr.Cmd, jackett.Cmd, cfg.Config)

	switch utils.IsExist(config.ConfPath) {
	case false:
		if config.ConfPath != config.DEFAULT_CONFIG_PATH {
			log.Fatalf("%s is not exists. Please check the file.", config.ConfPath)
		}

		if len(os.Args[1:]) == 0 {
			if err := RootCmd.Help(); err != nil {
				log.Fatalln(err)
			}
		}

	case true:
		config.ParseConfigFile()
	}
}
