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
