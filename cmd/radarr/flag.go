package radarr

import (
	"fmt"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "radarr",
	Short: "Radarr commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Flag("config").Value)
	},
}

func init() {
	// url := Cmd.Flags().StringP("url", "u", "http://localhost:7878", "Radarr URL")
	// api := Cmd.Flags().StringP("api", "a", "", "API Key")
	//
	// if *url != "" {
	// 	matched := utils.SplitUrl(*url)
	//
	// }
}
