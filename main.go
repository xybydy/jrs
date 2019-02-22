// Jackett, Radarr and Sonarr API wrapper is built to make the applications accessible via CLI easily.
// This project aims to create a base for automation mostly used commands for personal use cases.

// Current commands are:
// Jackett - addall: Command to add all public indexers.
// Jackett - get:	 Command to see all configured indexers.
// Jackett - list:	 Command to list all indexers available to use.
//
// Radarr - test: 	 Manually test all indexers to make them available for searching.
// Radarr - delete:  Delete all indexers added to Radarr.
// Radarr - add:	 Adds all configured indexers in jackett.
//
// Sonarr - test: 	 Manually test all indexers to make them available for searching.
// Sonarr - delete:  Delete all indexers added to Sonarr.
// Sonarr - add:	 Adds all configured indexers in jackett.
package main

import (
	"jrs/cmd/root"
	"log"
)

func main() {
	err := root.RootCmd.Execute()
	if err != nil {
		log.Fatalf("%s", err)
	}
}
