package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jrs/api"
	"jrs/config"
	"log"
	"net/http"
)

type App struct {
	client   *http.Client
	indexers api.Indexers
	jackett  *api.Jackett
}

func NewApp(conf *config.Config) *App {
	app := App{}
	app.jackett = api.NewJackett(conf)
	return &app

}

func (a *App) GetAllIndexers() {
	req := a.jackett.GetAllIndexers()
	resp, _ := a.client.Do(req)

	defer resp.Body.Close()

	msg, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(msg, &a.indexers)

}

func (a *App) getIndexerConfig(i *api.Indexer) ([]byte, error) {

	resp, err := a.client.Do(a.jackett.GetIndexerConfig(i.Id))
	if err != nil {
		log.Printf("%s\n", err)
	}
	defer resp.Body.Close()

	msg, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("%s\n", err)
	}

	return msg, err

}

func (a *App) AddIndexer(id, user, passwd string) {
	if len(a.indexers) == 0 {
		a.GetAllIndexers()
	}

	if index := a.indexers.GetIndexer(id); index != nil {
		if msg, err := a.getIndexerConfig(index); err == nil {
			var conf api.IndexerConfig
			json.Unmarshal(msg, &conf)
			conf.SetCredentials(user, passwd)

			mrs, err := json.Marshal(conf)
			if err != nil {
				log.Printf("q: %s\n", err)
			}

			resp, err := a.client.Do(a.jackett.UpdateIndexerConfig(id, string(mrs)))
			if err != nil {
				log.Printf("q: %s\n", err)
			}
			defer resp.Body.Close()

		} else {
			log.Printf("%s\n", err)
		}
	} else {
		log.Printf("There is no such indexer %s", id)
	}

}

func (a *App) AddAllPublicIndexers() {
	if len(a.indexers) == 0 {
		a.GetAllIndexers()
	}

	for _, i := range a.indexers {

		if i.Type == "public" && i.Configured == false {
			fmt.Printf("Posting %s\n", i.Name)

			msg, err := a.getIndexerConfig(&i)
			if err != nil {
				fmt.Printf("%s\n", err)
			}

			resp, err := a.client.Do(a.jackett.UpdateIndexerConfig(i.Id, string(msg)))
			if err != nil {
				fmt.Printf("q: %s\n", err)
			}
			defer resp.Body.Close()

			msg, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%s\n", err)
			}

		}
	}
}
