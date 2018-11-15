package jackett

import (
	"fmt"
	"net/http"
	"strings"

	"jrs/config"
	"net/http/cookiejar"
)

var (
	version = "2.0"
	root    = "/api"
)

type Jackett struct {
	version  string
	root     string
	api      string
	path     string
	indexers Indexers
	headers  http.Header
	client   *http.Client
}

func New(conf *config.Config) *Jackett {
	c := conf.GetDestination("jackett")
	j := &Jackett{version: version, root: root, api: c.Api, path: c.Path, headers: make(http.Header)}
	j.headers.Add("Content-Type", "application/json")
	j.client = new(http.Client)
	j.client.Jar, _ = cookiejar.New(nil)
	return j
}

func (j *Jackett) getAPIPath(category, action string, args ...string) string {
	path := j.path + "/api/v" + j.version + "/" + category

	if action != "" {
		path = path + "/" + action
	}

	return path
}

func (j *Jackett) ExportTorznab(indexerID string) string {
	return fmt.Sprintf(j.path + "/api/v" + j.version + "/indexers/" + indexerID + "/results/torznab/")
}

func (j *Jackett) ExportPotato(indexerID string) string {
	return fmt.Sprintf(j.path + "api/v" + j.version + "/indexers/" + indexerID + "/results/potato/")
}

func (j *Jackett) GetAPI() string {
	return j.api
}

func (j *Jackett) getAllIndexers() *http.Request {
	path := j.getAPIPath("indexers", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}

	return nil
}

func (j *Jackett) getServerConfig() *http.Request {
	path := j.getAPIPath("server", "config")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) getIndexerConfigReq(indexerID string) *http.Request {
	path := j.getAPIPath("indexers", indexerID+"/config")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) updateIndexerConfig(indexerID, config string) *http.Request {
	path := j.getAPIPath("indexers", indexerID+"/config")

	if req, err := http.NewRequest("POST", path, strings.NewReader(config)); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) deleteIndexer(indexerID string) *http.Request {
	path := j.getAPIPath("indexers", indexerID)

	if req, err := http.NewRequest("DELETE", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) testIndexer(indexerID string) *http.Request {
	path := j.getAPIPath("indexers", indexerID+"/test")
	if req, err := http.NewRequest("POST", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) resultsForIndexer(indexerID, query string) *http.Request {
	path := j.getAPIPath("indexers", indexerID+"/results?apikey="+j.api)

	if req, err := http.NewRequest("GET", path, strings.NewReader(query)); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) getServerCache() *http.Request {
	path := j.getAPIPath("server", "cache")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) getServerLogs() *http.Request {
	path := j.getAPIPath("server", "logs")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers

		return req
	}
	return nil
}

func (j *Jackett) updateServerConfig(serverConfig string) *http.Request {
	path := j.getAPIPath("server", "config")
	if req, err := http.NewRequest("POST", path, strings.NewReader(serverConfig)); err != nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) updateServer() *http.Request {
	path := j.getAPIPath("server", "update")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}

func (j *Jackett) updateAdminPassword(password string) *http.Request {
	path := j.getAPIPath("server", "adminpassword")

	if req, err := http.NewRequest("POST", path, strings.NewReader(password)); err == nil {
		req.Header = j.headers
		return req
	}
	return nil
}
