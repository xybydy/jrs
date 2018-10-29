package sonarr

import (
	"fmt"
	"io"
	"jrs/config"
	"jrs/utils"
	"net/http"
	"net/url"
	"strings"
)

type Sonarr struct {
	api     string
	path    string
	headers http.Header
}

func New(c *config.Config) *Sonarr {
	conf := c.GetDestination("Sonarr")
	s := &Sonarr{conf.API, utils.BuildURL(conf.IP, conf.Port), http.Header{}}
	s.headers.Add("Content-Type", "application/json")
	s.headers.Add("X-Api-Key", s.api)
	return s
}

func (s *Sonarr) GetHeaders() http.Header {
	return s.headers
}

func (s *Sonarr) MakeRequest(method string, body io.Reader, args ...string) (*http.Request, error) {
	// TODO URL Base olmayinca cogu komut calismiyor. URL base eklemek gerekecek.
	path := s.path + "/api"

	for _, arg := range args {
		path = path + "/" + arg
	}

	if request, err := http.NewRequest(method, path, body); err == nil {
		request.Header = s.headers
		return request, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) Calendar() (*http.Request, error) {
	// path := s.getAPIPath("calendar", "")
	if req, err := s.MakeRequest("GET", nil, "calendar"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetOngoingCommands() (*http.Request, error) {
	// path := s.getAPIPath("command", "")
	if req, err := s.MakeRequest("GET", nil, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetCommandStatus(id string) (*http.Request, error) {
	// path := s.getAPIPath("command", id)
	if req, err := s.MakeRequest("GET", nil, "command", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) RefreshSeries(id string) (*http.Request, error) {
	data := url.Values{}
	data.Add("name", "refreshseries")

	if id != "" {
		data.Add("seriesId", id)
	}

	// path := s.getAPIPath("command", "")

	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) RescanSeries(id string) (*http.Request, error) {
	data := url.Values{}
	data.Add("name", "rescanseries")

	if id != "" {
		data.Add("seriesId", id)
	}

	// path := s.getAPIPath("command", "")

	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) EpisodeSearch(ids []int) (*http.Request, error) {
	data := url.Values{}
	data.Add("name", "episodesearch")

	if len(ids) != 0 || ids != nil {
		for _, id := range ids {
			data.Add("episodeIds", string(id))
		}
	}

	// path := s.getAPIPath("command", "")
	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) SeasonSearch(seriesID, seasonNumber string) (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "seasonsearch")
	data.Add("seriesId", seriesID)
	data.Add("seasonNumber", seasonNumber)

	// path := s.getAPIPath("command", "")
	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) SeriesSearch(seriesID string) (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "seriessearch")
	data.Add("seriesId", seriesID)

	// path := s.getAPIPath("command", "")
	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) RssSync() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "rsssync")

	// path := s.getAPIPath("command", "")
	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) RenameFiles(files []int) (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "renamefiles")

	if len(files) != 0 || files != nil {
		for _, id := range files {
			data.Add("files", string(id))
		}
	}

	// path := s.getAPIPath("command", "")
	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) RenameSeries(seriesIds []int) (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "renameseries")

	if len(seriesIds) != 0 || seriesIds != nil {
		for _, id := range seriesIds {
			data.Add("seriesIds", string(id))
		}
	}

	// path := s.getAPIPath("command", "")
	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) DiskSpace() (*http.Request, error) {
	// path := s.getAPIPath("diskspace", "")
	if req, err := s.MakeRequest("POST", nil, "diskspace"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) Backup() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "backup")

	// path := s.getAPIPath("command", "")

	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) missingEpisodeSearch() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "missingEpisodeSearch")

	// path := s.getAPIPath("command", "")

	if req, err := s.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetShow(seriesID string) (*http.Request, error) {
	data := url.Values{}

	data.Add("seriesId", seriesID)

	// path := s.getAPIPath("episode", "")

	if req, err := s.MakeRequest("GET", strings.NewReader(data.Encode()), "episode"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) GetEpisode(episodeID string) (*http.Request, error) {
	// path := s.getAPIPath("episode", episodeID)

	if req, err := s.MakeRequest("GET", nil, "episode", episodeID); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// TODO Episode PUT to be implemented

func (s *Sonarr) GetEpisodeFiles(seriesID string) (*http.Request, error) {
	data := url.Values{}

	data.Add("seriesId", seriesID)

	// path := s.getAPIPath("episodefile", "")

	if req, err := s.MakeRequest("GET", strings.NewReader(data.Encode()), "episodefile"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) GetShowFiles(id string) (*http.Request, error) {
	// path := s.getAPIPath("episodefile", id)

	if req, err := s.MakeRequest("GET", nil, "episodefile", id); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) DeleteEpisode(id string) (*http.Request, error) {
	// path := s.getAPIPath("episodefile", id)

	if req, err := s.MakeRequest("DELETE", nil, "episodefile", id); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) History(sortKey, page, pageSize, sortDir string) (*http.Request, error) {
	data := url.Values{}

	if sortKey == "" {
		data.Add("sortKey", "date") // series.title OR data
	} else {
		data.Add("sortKey", sortKey)
	}

	if page != "" {
		data.Add("page", page)
	}

	if pageSize != "" {
		data.Add("pageSize", pageSize)
	}

	if sortDir != "" {
		data.Add("sortDir", sortDir)
	}

	// path := s.getAPIPath("history", "")

	if req, err := s.MakeRequest("GET", nil, "history"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) WantedMissing(sortKey, page, pageSize, sortDir string) (*http.Request, error) {
	data := url.Values{}

	if sortKey == "" {
		data.Add("sortKey", "airDateUtc") // series.title OR airDateUtc
	} else {
		data.Add("sortKey", sortKey)
	}

	if page != "" {
		data.Add("page", page)
	}

	if pageSize != "" {
		data.Add("pageSize", pageSize)
	}

	if sortDir != "" {
		data.Add("sortDir", sortDir)
	}

	// path := s.getAPIPath("wanted", "missing")

	if req, err := s.MakeRequest("GET", nil, "wanted", "missing"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) Queue() (*http.Request, error) {
	// path := s.getAPIPath("queue", "")

	if req, err := s.MakeRequest("GET", nil, "queue"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) RemoveDownload(id string, blacklist bool) (*http.Request, error) {
	data := url.Values{}

	data.Add("id", id)

	if blacklist == true {
		data.Add("blacklist", "true")
	} else {
		data.Add("blacklist", "false")
	}

	// path := s.getAPIPath("queue", "")

	if req, err := s.MakeRequest("GET", strings.NewReader(data.Encode()), "queue"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) GetProfiles() (*http.Request, error) {
	// path := s.getAPIPath("profile", "")

	if req, err := s.MakeRequest("GET", nil, "profile"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) ConfigMediaManagement() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "config", "mediamanagement"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigNaming() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "config", "naming"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (s *Sonarr) ConfigIndexer() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "config", "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigDownloadClient() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "config", "downloadclient"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetDownloadClient() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "downloadclient"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigNotification() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "config", "notification"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigHost() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "config", "host"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigUI() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "config", "ui"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigNetImport() (*http.Request, error) {
	if req, err := s.MakeRequest("GET", nil, "config", "netimport"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetNetImport() (*http.Request, error) {
	if req, err := s.MakeRequest("GET", nil, "netimport"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigProfile() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "profile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigDelayProfile() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "delayprofile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) ConfigQualityDefinition() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "qualitydefinition"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetIndexers() (*http.Request, error) { // same
	if req, err := s.MakeRequest("GET", nil, "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetRemotePathMapping() (*http.Request, error) {
	if req, err := s.MakeRequest("GET", nil, "remotePathMapping"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetRootFolder() (*http.Request, error) {
	if req, err := s.MakeRequest("GET", nil, "rootfolder"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetExclusions() (*http.Request, error) {
	if req, err := s.MakeRequest("GET", nil, "exclusions"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetMetadata() (*http.Request, error) {
	if req, err := s.MakeRequest("GET", nil, "metadata"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) GetIndexerSchema() (*http.Request, error) {
	if req, err := s.MakeRequest("GET", nil, "indexer", "schema"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) DeleteIndexer(i IndexerSchema) (*http.Request, error) {
	id := fmt.Sprintf("%v", i.ID)
	if req, err := s.MakeRequest("DELETE", nil, "indexer", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (s *Sonarr) SetIndexer(i IndexerSchema) (*http.Request, error) {
	if req, err := s.MakeRequest("POST", nil, "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
