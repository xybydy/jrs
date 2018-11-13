package cmd

import (
	"github.com/spf13/cobra"
	cfg "jrs/cmd/config"
	"jrs/cmd/jackett"
	"jrs/cmd/radarr"
	"jrs/cmd/sonarr"
	"jrs/config"
)

var RootCmd = &cobra.Command{Use: "jrs",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&config.ConfPath, "config", "c", "config.toml", "Config file path")
	config.ParseConfigFile()

	RootCmd.AddCommand(radarr.Cmd)
	RootCmd.AddCommand(sonarr.Cmd)
	RootCmd.AddCommand(jackett.Cmd)
	RootCmd.AddCommand(cfg.Config)
}
