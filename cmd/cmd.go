// Terminal command functions and helper functions resides on this package
package cmd

import (
	"jrs/config"
	"log"
	"strings"
)

// Checks the parameters of the applications in the config file
// If parameters passed as command line parameters, it overrides the parameters in the config file.
// If either there is no parameters passed nor exist in config file error occurs.
func GetConfig(dest, url, api string) {
	app := config.Params.GetDestination(dest)

	if app.Path == "" && url == "" {
		log.Fatalf("There is path specified for %s. Please set config file or provide path with --path option.", dest)
	}
	if app.Api == "" && api == "" {
		log.Fatalf("There is api specified for %s. Please set config file or provide api with --api option.", api)
	}

	if url != "" {
		config.Params.ChangeParams(strings.Title(dest), "Path", url)
	}
	if api != "" {
		config.Params.ChangeParams(strings.Title(dest), "Api", api)
	}
}
