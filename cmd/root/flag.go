package root

import (
	cfg "jrs/cmd/config"
	"jrs/cmd/jackett"
	"jrs/cmd/radarr"
	"jrs/cmd/sonarr"
	"jrs/config"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{Use: "jrs",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&config.ConfPath, "config", "c", "config.toml", "Config file path")
	config.ParseConfigFile()
	RootCmd.AddCommand(radarr.Cmd, sonarr.Cmd, jackett.Cmd, cfg.Config)
}
