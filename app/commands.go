package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"jrs/api/jackett"
	"jrs/config"
	"log"
	"net/http"
	"net/http/cookiejar"
)

type App struct {
	client   http.Client
	jar      *cookiejar.Jar
	indexers jackett.Indexers
	Jackett  *jackett.Jackett
}

func NewApp(conf *config.Config) *App {
	app := App{}
	app.Jackett = jackett.NewJackett(conf)
	app.jar, _ = cookiejar.New(nil)
	app.client.Jar = app.jar
	return &app
}

func (a *App) GetAllIndexers() {
	req := a.Jackett.GetAllIndexers()
	resp, err := a.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	msg, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(msg, &a.indexers)
}

func (a *App) GetConfiguredIndexers() jackett.Indexers {
	inx := jackett.Indexers{}
	if len(a.indexers) == 0 {
		a.GetAllIndexers()
	}
	for _, i := range a.indexers {
		if i.Configured == true {
			inx = append(inx, i)
		}
	}
	return inx
}

func (a *App) getIndexerConfig(i *jackett.Indexer) ([]byte, error) {

	resp, err := a.client.Do(a.Jackett.GetIndexerConfig(i.ID))
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
			var conf jackett.IndexerConfig
			json.Unmarshal(msg, &conf)
			conf.SetCredentials(user, passwd)

			mrs, err := json.Marshal(conf)
			if err != nil {
				log.Printf("q: %s\n", err)
			}
			resp, err := a.client.Do(a.Jackett.UpdateIndexerConfig(id, string(mrs)))
			if err != nil {
				log.Printf("q: %s\n", err)
			}
			if resp.StatusCode != 200 {
				data := make(map[string]interface{})
				msg, err = ioutil.ReadAll(resp.Body)
				if err != nil {
					log.Fatal(err)
				}
				err := json.Unmarshal(msg, &data)
				if err != nil {
					log.Fatal(err)
				}
				log.Fatalf("%s - %s", index.Name, data["error"])
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

			resp, err := a.client.Do(a.Jackett.UpdateIndexerConfig(i.ID, string(msg)))
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
