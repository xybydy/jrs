package config

import (
	"github.com/spf13/cobra"
)

var (
	url string
	api string
)

var config = &cobra.Command{
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
}

func init() {
	radarr.AddCommand(save)
	sonarr.AddCommand(save)
	jackett.AddCommand(save)
	set.AddCommand(radarr, sonarr, jackett)
	config.AddCommand(set)
}
