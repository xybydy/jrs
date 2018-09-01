package main

import (
	"flag"
	"jrs/config"
	"log"
	"os"
	"strings"

	"net/http"

	"jrs/api/sonarr"

	"io/ioutil"

	"encoding/json"

	"jrs/app"

	"bytes"

	"fmt"

	"jrs/api/radarr"

	"github.com/burntsushi/toml"
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

	if params.Jackett.API == "" {
		param := os.Getenv("JACKETT")
		if param != "" {
			params.Jackett.API = param
		} else {
			log.Fatal("There is no Jackett API configured.")
		}
	}

	if len(params.Dest) < 1 {
		log.Fatal("There is no Destination configured.")
	}

	for _, dest := range params.Dest {
		if dest.API == "" {
			param := os.Getenv(strings.ToUpper(dest.Name))

			if param != "" {
				params.GetDestination(dest.Name).API = param
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

func SonarrTestAllIndexers() {
	var schemas sonarr.IndexerSchemas
	s := sonarr.New(params)

	client := new(http.Client)

	req, _ := s.GetIndexers()
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	json.Unmarshal(body, &schemas)

	for _, i := range schemas {
		data, err := json.Marshal(i)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		req, err := s.MakeRequest("POST", bytes.NewBuffer(data), "indexer", "test")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		resp, _ := client.Do(req)

		fmt.Printf("%v - %v\n", i.Name, resp.StatusCode)

	}

}

// TODO change api code for all indexer in R and S
func SonarrAddAllIndexes() {
	var schema sonarr.IndexerSchemas
	s := sonarr.New(params)
	j := app.NewApp(params)
	inx := j.GetConfiguredIndexers()

	// Grabbed schema
	schm, err := s.GetIndexerSchema()
	if err != nil {
		log.Fatalf("%v", err)
	}
	client := new(http.Client)
	resp, err := client.Do(schm)

	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	json.Unmarshal(body, &schema)

	torznab := schema.GetTorznab()

	for _, indexer := range inx {
		torznab.Name = indexer.Name
		torznab.EnableRss = true
		torznab.EnableSearch = true
		torznab.SupportRss = true
		torznab.SupportSearch = true
		for i := range torznab.Fields {
			if torznab.Fields[i].Name == "BaseUrl" {
				torznab.Fields[i].Value = j.Jackett.ExportTorznab(indexer.ID)
			}
			if torznab.Fields[i].Name == "ApiKey" {
				torznab.Fields[i].Value = j.Jackett.GetAPI()
			}

		}
		data, err := json.Marshal(torznab)
		if err != nil {
			log.Fatalf("s", err)
		}
		req, _ := s.MakeRequest("POST", bytes.NewBuffer(data), "indexer")

		resp, err := client.Do(req)
		if err != nil {
			log.Print("HATA", err)
		}

		log.Printf("%v - %v\n", indexer.Name, resp.StatusCode)

		// f, _ := ioutil.ReadAll(resp.Body)
		// fmt.Printf("%s\n", f)
		// resp.Body.Close()

	}
}

func SonarrDeleteAllIndexers() {
	var schemas sonarr.IndexerSchemas
	s := sonarr.New(params)
	client := new(http.Client)

	req, err := s.GetIndexers()
	if err != nil {
		fmt.Printf("%v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	json.Unmarshal(data, &schemas)

	for _, i := range schemas {
		req, _ = s.DeleteIndexer(i)
		client.Do(req)

	}
	resp.Body.Close()

}

func RadarrTestAllIndexers() {
	var schemas radarr.IndexerSchemas
	r := radarr.New(params)

	client := new(http.Client)

	req, _ := r.GetIndexers()
	resp, _ := client.Do(req)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}

	json.Unmarshal(body, &schemas)

	for _, i := range schemas {
		data, err := json.Marshal(i)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		req, err := r.MakeRequest("POST", bytes.NewBuffer(data), "indexer", "test")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		resp, _ := client.Do(req)

		fmt.Printf("%v - %v\n", i.Name, resp.StatusCode)

	}

}

// TODO change api code for all indexer in R and S
func RadarrAddAllIndexes() {
	var schema radarr.IndexerSchemas
	r := radarr.New(params)
	j := app.NewApp(params)
	inx := j.GetConfiguredIndexers()

	// Grabbed schema
	schm, err := r.GetIndexerSchema()
	if err != nil {
		log.Fatalf("1 - %v", err)
	}
	client := new(http.Client)
	resp, err := client.Do(schm)

	if err != nil {
		log.Fatalf("%v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	json.Unmarshal(body, &schema)

	torznab := schema.GetTorznab()

	for _, indexer := range inx {
		torznab.Name = indexer.Name
		torznab.EnableRss = true
		torznab.EnableSearch = true
		torznab.SupportRss = true
		torznab.SupportSearch = true
		for i := range torznab.Fields {
			if torznab.Fields[i].Name == "BaseUrl" {
				torznab.Fields[i].Value = j.Jackett.ExportTorznab(indexer.ID)
			}
			if torznab.Fields[i].Name == "ApiKey" {
				torznab.Fields[i].Value = j.Jackett.GetAPI()
			}
			if torznab.Fields[i].Name == "RequiredFlags" {
				torznab.Fields[i].Value = ""
			}

		}
		data, err := json.Marshal(torznab)
		if err != nil {
			log.Fatalf("s", err)
		}
		req, _ := r.MakeRequest("POST", bytes.NewBuffer(data), "indexer")

		resp, err := client.Do(req)
		if err != nil {
			log.Print("HATA", err)
		}

		log.Printf("%v - %v\n", indexer.Name, resp.StatusCode)

		// f, _ := ioutil.ReadAll(resp.Body)
		// fmt.Printf("%s\n", f)
		// resp.Body.Close()

	}
}

func RadarrDeleteAllIndexers() {
	var schemas radarr.IndexerSchemas
	r := radarr.New(params)
	client := new(http.Client)

	req, err := r.GetIndexers()
	if err != nil {
		fmt.Printf("%v", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("%v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	json.Unmarshal(data, &schemas)

	for _, i := range schemas {
		req, _ = r.DeleteIndexer(i)
		client.Do(req)

	}
	resp.Body.Close()

}

func main() {
	// RadarrDeleteAllIndexers()
	RadarrAddAllIndexes()
	// RadarrTestAllIndexers()
	// SonarrTestAllIndexers()

}

// TODO http response 303 to be fixed
