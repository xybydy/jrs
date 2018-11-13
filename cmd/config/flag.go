package config

import (
	"fmt"
	"github.com/spf13/cobra"
	"jrs/config"
)

var (
	url string
	api string
)

var Config = &cobra.Command{
	Use: "config",
}

var set = &cobra.Command{
	Use: "set",
}

var radarr = &cobra.Command{
	Use: "radarr",
}

var sonarr = &cobra.Command{
	Use: "sonarr",
}

var jackett = &cobra.Command{
	Use: "jackett",
}

var save = &cobra.Command{
	Use: "save",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args[0])
		config.Params.SaveFile(args[0])
	},
}

func init() {
	radarr.AddCommand(save)
	sonarr.AddCommand(save)
	jackett.AddCommand(save)
	set.AddCommand(radarr, sonarr, jackett)
	Config.AddCommand(set)
}
