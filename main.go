package main

import (
	"flag"
	"jrs/config"
	"log"
	"os"
	"strings"

	"fmt"
	"github.com/burntsushi/toml"
	"io/ioutil"
	"net/http"
)

var (
	params   *config.Config
	confArgs string
)

func init() {
	flag.StringVar(&confArgs, "config", "config.toml", "Config File")
	flag.Parse()

	if _, err := toml.DecodeFile(confArgs, &params); err != nil {
		log.Fatal(err)
	}

	if params.Jackett.Api == "" {
		param := os.Getenv("JACKETT")
		if param != "" {
			params.Jackett.Api = param
		} else {
			log.Fatal("There is no Jackett API configured.")
		}
	}

	if len(params.Dest) < 1 {
		log.Fatal("There is no Destination configured.")
	}

	for _, dest := range params.Dest {
		if dest.Api == "" {
			param := os.Getenv(strings.ToUpper(dest.Name))

			if param != "" {
				params.GetDestination(dest.Name).Api = param
			} else {
				if len(params.Dest) == 1 {
					log.Fatalf("There is no %s API configured.", dest.Name)
				} else {
					log.Printf("There is no %s API configured.", dest.Name)
				}
			}
		}

	}
}

func main() {
	req, _ := http.NewRequest("GET", "http://192.168.1.21:9093/api/seriesId?apikey=18573072d90e4752b80fddd247fb80f1", nil)
	client := new(http.Client)
	q, _ := client.Do(req)
	defer q.Body.Close()
	ee, _ := ioutil.ReadAll(q.Body)
	fmt.Printf("%s", ee)
	// TODO http response 303 to be fixed

}
