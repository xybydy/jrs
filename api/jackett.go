package api

import (
	"fmt"
	"jrs/config"
	"jrs/utils"
	"net/http"
	"strings"
)

var (
	version = "2.0"
	root    = "/api"
)

type Caps struct {
	Id   string `json:"ID"`
	Name string `json:"Name"`
}

type Indexer struct {
	Id               string
	Name             string
	Description      string
	Type             string
	Configured       bool
	SiteLink         string `json:"site_link"`
	Alternativelinks []string
	Language         string
	LastError        string `json:"last_error"`
	Potatoenabled    bool
	Caps             []Caps
}

type IndexerConfig []struct {
	Id      string
	Type    string            `json:,omitempty`
	Name    string            `json:,omitempty`
	Value   string            `json:,omitempty`
	Options map[string]string `json:,omitempty`
}

func (ic *IndexerConfig) UpdateField(id, param string) {
	for _, i := range *ic {
		if i.Id == id {
			i.Value = param
		}
	}
}

func (ic *IndexerConfig) SetCredentials(user, passwd string) {
	ic.UpdateField("username", user)
	ic.UpdateField("password", passwd)
}

type Indexers []Indexer

func (i *Indexers) GetIndexer(id string) *Indexer {
	for _, k := range *i {
		if k.Id == id {
			return &k
		}
	}
	return nil
}

type Jackett struct {
	version string
	root    string
	api     string
	path    string
	headers http.Header
}

func NewJackett(conf *config.Config) *Jackett {
	j := &Jackett{version: version, root: root, api: conf.Jackett.Api, path: utils.BuildURL(conf.Jackett.Ip, conf.Jackett.Port)}
	j.headers.Add("Content-Type", "application/json")
	return j
}

func (j *Jackett) getApiPath(category, action string, args ...string) string {
	path := j.path + "/api/v" + j.version + "/" + category

	if action != "" {
		path = path + "/" + action
	}

	return path
}

func (j *Jackett) ExportTorznab(indexerId string) string {
	return fmt.Sprintf(j.path + "api/v" + j.version + "/indexers/" + indexerId + "/results/torznab/")
}

func (j *Jackett) ExportPotato(indexerId string) string {
	return fmt.Sprintf(j.path + "api/v" + j.version + "/indexers/" + indexerId + "/results/potato/")
}

func (j *Jackett) GetAPI() string {
	return j.api
}

func (j *Jackett) GetAllIndexers() *http.Request {
	path := j.getApiPath("indexers", "")
	fmt.Println(path)
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}

	return nil
}

func (j *Jackett) GetServerConfig() *http.Request {
	path := j.getApiPath("server", "config")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) GetIndexerConfig(indexerId string) *http.Request {
	path := j.getApiPath("indexers", indexerId+"/config")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) UpdateIndexerConfig(indexerId, config string) *http.Request {
	path := j.getApiPath("indexers", indexerId+"/config")

	if req, err := http.NewRequest("POST", path, strings.NewReader(config)); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) DeleteIndexer(indexerId string) *http.Request {
	path := j.getApiPath("indexers", indexerId)

	if req, err := http.NewRequest("DELETE", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) TestIndexer(indexerId string) *http.Request {
	path := j.getApiPath("indexers", indexerId+"/test")
	if req, err := http.NewRequest("POST", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) ResultsForIndexer(indexerId, query string) *http.Request {
	path := j.getApiPath("indexers", indexerId+"/results?apikey="+j.api)

	if req, err := http.NewRequest("GET", path, strings.NewReader(query)); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) GetServerCache() *http.Request {
	path := j.getApiPath("server", "cache")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) GetServerLogs() *http.Request {
	path := j.getApiPath("server", "logs")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers

		return req
	}
	return nil
}

func (j *Jackett) UpdateServerConfig(serverConfig string) *http.Request {
	path := j.getApiPath("server", "config")
	if req, err := http.NewRequest("POST", path, strings.NewReader(serverConfig)); err != nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) UpdateServer() *http.Request {
	path := j.getApiPath("server", "update")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) UpdateAdminPassword(password string) *http.Request {
	path := j.getApiPath("server", "adminpassword")

	if req, err := http.NewRequest("POST", path, strings.NewReader(password)); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}
