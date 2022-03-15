package arr

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/xybydy/jrs/jackett"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Arr struct {
	ApiKey  string
	ApiVer  string
	URL     string
	User    string
	Pass    string
	Client  *http.Client
	Timeout time.Duration
}

func (a *Arr) Request(method string, body io.Reader, args ...string) (*http.Request, error) {
	path := a.URL + "/api" + "/" + a.ApiVer

	for _, arg := range args {
		path = path + "/" + arg
	}

	ctx, cancel := context.WithTimeout(context.Background(), a.Timeout)
	defer cancel()

	if request, err := http.NewRequestWithContext(ctx, method, path, body); err == nil {
		request.Header.Add("Content-Type", "application/json")
		request.Header.Add("X-Api-Key", a.ApiKey)
		return request, err
	} else {
		return nil, err
	}
}

// bunu method yerine duz func yapalim
func (a *Arr) TestAllIndexers() {
	var schemas IndexerSchemas

	req, _ := a.GetIndexers()
	resp, _ := a.Client.Do(req)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer resp.Body.Close()
	log.Print(req, body)
	err = json.Unmarshal(body, &schemas)
	if err != nil {
		fmt.Printf("%v", err)
	}

	for _, i := range schemas {
		data, err := json.Marshal(i)
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		req, err := a.Request(http.MethodPost, bytes.NewBuffer(data), "indexer", "test")
		if err != nil {
			log.Fatalf("%v\n", err)
		}
		resp, _ := a.Client.Do(req)

		fmt.Printf("%v - %v\n", i.Name, resp.StatusCode)
	}
}

// bunu method yerine duz func yapalim
func (a *Arr) AddAllIndexers(j *jackett.Jackett, hostdomain string) {
	var schema IndexerSchemas

	inx := j.GetConfiguredIndexers()
	schm, err := a.GetIndexerSchema()
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
		req, _ := a.Request(http.MethodPost, bytes.NewBuffer(data), "indexer")

		resp, err := client.Do(req)
		if err != nil {
			log.Print("HATA", err)
		}
		log.Printf("%v - %v\n", indexer.Name, resp.StatusCode)
	}

}

func (a *Arr) DeleteAllIndexers() {
	var schemas IndexerSchemas

	req, err := a.GetIndexers()
	if err != nil {
		fmt.Printf("%v", err)
	}

	resp, err := a.Client.Do(req)
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
		req, _ = a.DeleteIndexer(i)
		_, err = a.Client.Do(req)
		if err != nil {
			fmt.Printf("%v", err)
		}
	}
	resp.Body.Close()
}
