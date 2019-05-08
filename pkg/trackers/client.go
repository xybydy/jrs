package trackers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jrs/pkg/jackett"
	"log"
	"net/http"
)

// Common methods of applications resides here
type CommonTracker struct {
	Client  *http.Client
	Api     string
	Path    string
	Headers http.Header
}

func CreateTracker(t Tracker, api, path string) *CommonTracker {
	return &CommonTracker{Client: new(http.Client), Headers: http.Header{}, Api: api, Path: path}
}

func (c *CommonTracker) TestAllIndexers() {
	var schemas IndexerSchemas

	req, _ := c.t.GetIndexers()
	resp, _ := c.Client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()

	err = json.Unmarshal(body, &schemas)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for _, i := range schemas {
		data, err := json.Marshal(i)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		req, err := c.t.BuildRequest("POST", bytes.NewBuffer(data), "indexer", "test")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		resp, _ := c.Client.Do(req)

		fmt.Printf("%v - %v\n", i.Name, resp.StatusCode)
	}
}

func (c *CommonTracker) AddAllIndexers(j *jackett.Jackett) {
	var schema IndexerSchemas

	inx := j.GetConfiguredIndexers()
	schm, err := c.t.GetIndexerSchema()
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
	if err = json.Unmarshal(body, &schema); err != nil {
		log.Fatalf("%v", err)
	}

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

		}
		data, err := json.Marshal(torznab)
		if err != nil {
			log.Fatalf("%s", err)
		}
		req, _ := c.t.BuildRequest("POST", bytes.NewBuffer(data), "indexer")

		resp, err := client.Do(req)
		if err != nil {
			log.Print("HATA", err)
		}
		log.Printf("%v - %v\n", indexer.Name, resp.StatusCode)
	}

}

func (c *CommonTracker) DeleteAllIndexers() {
	var schemas IndexerSchemas

	req, err := c.t.GetIndexers()
	if err != nil {
		fmt.Printf("%v", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		fmt.Printf("%v", err)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%v", err)
	}
	if err = json.Unmarshal(data, &schemas); err != nil {
		log.Fatalf("%v", err)
	}

	for _, i := range schemas {
		req, _ = c.t.DeleteIndexer(i)
		_, err = c.Client.Do(req)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	resp.Body.Close()
}
