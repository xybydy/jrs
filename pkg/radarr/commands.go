package radarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jrs/config"
	"jrs/pkg/jackett"
	"log"
	"net/http"
)

type client struct {
	r      *Radarr
	client *http.Client
}

func NewClient() *client {
	c := new(client)
	c.r = New(config.Params)
	return c
}

func (c *client) TestAllIndexers() {
	var schemas IndexerSchemas

	req, _ := c.r.GetIndexers()
	resp, _ := c.client.Do(req)

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
		req, err := c.r.BuildRequest("POST", bytes.NewBuffer(data), "indexer", "test")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		resp, _ := c.client.Do(req)

		fmt.Printf("%v - %v\n", i.Name, resp.StatusCode)
	}

}

func (c *client) AddAllIndexers() {
	var schema IndexerSchemas

	j := jackett.New(config.Params)

	inx := j.GetConfiguredIndexers()
	schm, err := c.r.GetIndexerSchema()
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
				torznab.Fields[i].Value = j.ExportTorznab(indexer.ID)
			}
			if torznab.Fields[i].Name == "ApiKey" {
				torznab.Fields[i].Value = j.GetAPI()
			}
			if torznab.Fields[i].Name == "RequiredFlags" {
				torznab.Fields[i].Value = ""
			}

		}
		data, err := json.Marshal(torznab)
		if err != nil {
			log.Fatalf("s", err)
		}
		req, _ := c.r.BuildRequest("POST", bytes.NewBuffer(data), "indexer")

		resp, err := client.Do(req)
		if err != nil {
			log.Print("HATA", err)
		}

		log.Printf("%v - %v\n", indexer.Name, resp.StatusCode)

	}

}

func (c *client) DeleteAllIndexers() {
	var schemas IndexerSchemas

	req, err := c.r.GetIndexers()
	if err != nil {
		fmt.Printf("%v", err)
	}

	resp, err := c.client.Do(req)
	if err != nil {
		fmt.Printf("%v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	json.Unmarshal(data, &schemas)

	for _, i := range schemas {
		req, _ = c.r.DeleteIndexer(i)
		c.client.Do(req)
	}
	resp.Body.Close()
}
