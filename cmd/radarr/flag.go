package radarr

import (
	"flag"
	"github.com/spf13/cobra"
)

func ParseOptions() {
	var configPath string
	flag.StringVar(&configPath, "config", "config.toml", "Config file path")

}

var RadarrCmd = &cobra.Command{
	Use:   "radarr",
	Short: "Radarr commands",
}

var RootCmd = &cobra.Command{Use: "jrs"}

func init() {
	RootCmd.AddCommand(RadarrCmd)
}
