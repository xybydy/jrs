package trackers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xybydy/jrs/pkg/jackett"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	C      Tracker
	Client *http.Client
}

func (c *Client) TestAllIndexers() {
	var schemas IndexerSchemas

	req, _ := c.C.GetIndexers()
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
		req, err := c.C.BuildRequest("POST", bytes.NewBuffer(data), "indexer", "test")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		resp, _ := c.Client.Do(req)

		fmt.Printf("%v - %v\n", i.Name, resp.StatusCode)
	}
}

func (c *Client) AddAllIndexers(j *jackett.Jackett, hostdomain string) {
	var schema IndexerSchemas

	inx := j.GetConfiguredIndexers()
	schm, err := c.C.GetIndexerSchema()
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
				torznab.Fields[i].Value = j.ExportTorznab(indexer.ID, hostdomain)
			}
			if torznab.Fields[i].Name == "ApiKey" {
				torznab.Fields[i].Value = j.GetAPI()
			}

		}
		data, err := json.Marshal(torznab)
		if err != nil {
			log.Fatalf("%s", err)
		}
		req, _ := c.C.BuildRequest("POST", bytes.NewBuffer(data), "indexer")

		resp, err := client.Do(req)
		if err != nil {
			log.Print("HATA", err)
		}
		log.Printf("%v - %v\n", indexer.Name, resp.StatusCode)
	}

}

func (c *Client) DeleteAllIndexers() {
	var schemas IndexerSchemas

	req, err := c.C.GetIndexers()
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
		req, _ = c.C.DeleteIndexer(i)
		_, err = c.Client.Do(req)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	resp.Body.Close()
}
