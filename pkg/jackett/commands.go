package jackett

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func (j *Jackett) GetAllIndexers() {
	req := j.getAllIndexers()
	resp, err := j.client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	msg, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(msg, &j.indexers)
	if err != nil {
		log.Fatalf("%s", err)
	}
}

func (j *Jackett) GetConfiguredIndexers() Indexers {
	inx := Indexers{}
	if len(j.indexers) == 0 {
		j.GetAllIndexers()
	}
	for _, i := range j.indexers {
		if i.Configured {
			inx = append(inx, i)
		}
	}
	return inx
}

func (j *Jackett) getIndexerConfig(i *Indexer) ([]byte, error) {

	resp, err := j.client.Do(j.getIndexerConfigReq(i.ID))
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

func (j *Jackett) AddIndexer(id, user, passwd string) {
	if len(j.indexers) == 0 {
		j.GetAllIndexers()
	}

	if index := j.indexers.GetIndexer(id); index != nil {
		if msg, err := j.getIndexerConfig(index); err == nil {
			var conf IndexerConfig
			err = json.Unmarshal(msg, &conf)
			if err != nil {
				log.Fatalf("%s", err)
			}
			conf.SetCredentials(user, passwd)

			mrs, err := json.Marshal(conf)
			if err != nil {
				log.Printf("q: %s\n", err)
			}
			resp, err := j.client.Do(j.updateIndexerConfig(id, string(mrs)))
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

func (j *Jackett) AddAllPublicIndexers() {
	if len(j.indexers) == 0 {
		j.GetAllIndexers()
	}

	for _, i := range j.indexers {

		if i.Type == "public" && !i.Configured {
			fmt.Printf("Posting %s\n", i.Name)

			msg, err := j.getIndexerConfig(&i)
			if err != nil {
				fmt.Printf("%s\n", err)
			}

			resp, err := j.client.Do(j.updateIndexerConfig(i.ID, string(msg)))
			if err != nil {
				fmt.Printf("q: %s\n", err)
			}
			defer resp.Body.Close()

			_, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("%s\n", err)
			}

		}
	}
}
