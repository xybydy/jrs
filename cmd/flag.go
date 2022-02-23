package cmd

import (
	"github.com/xybydy/jrs/config"
	"log"
	"strings"
)

func CheckConfig(dest, url, api string) {
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
