package api

import (
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

func NewSonarr(c *config.Config) *Sonarr {
	conf := c.GetDestination("Sonarr")
	s := &Sonarr{conf.Api, utils.BuildURL(conf.Ip, conf.Port), http.Header{}}
	s.headers.Add("Content-Type", "application/json")
	s.headers.Add("X-Api-Key", s.api)
	return s
}

func (s *Sonarr) getAPIPath(args ...string) string {
	path := s.path + "/api"

	for _, arg := range args {
		path = path + "/" + arg
	}

	return path
}

func (s *Sonarr) Calendar() *http.Request {
	path := s.getAPIPath("calendar", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) GetOngoingCommands() *http.Request {
	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) GetCommandStatus(id string) *http.Request {
	path := s.getAPIPath("command", id)
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) RefreshSeries(id string) *http.Request {
	data := url.Values{}
	data.Add("name", "refreshseries")

	if id != "" {
		data.Add("seriesId", id)
	}

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) RescanSeries(id string) *http.Request {
	data := url.Values{}
	data.Add("name", "rescanseries")

	if id != "" {
		data.Add("seriesId", id)
	}

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) EpisodeSearch(ids []int) *http.Request {
	data := url.Values{}
	data.Add("name", "episodesearch")

	if len(ids) != 0 || ids != nil {
		for _, id := range ids {
			data.Add("episodeIds", string(id))
		}
	}

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) SeasonSearch(seriesID, seasonNumber string) *http.Request {
	data := url.Values{}

	data.Add("name", "seasonsearch")
	data.Add("seriesId", seriesID)
	data.Add("seasonNumber", seasonNumber)

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) SeriesSearch(seriesID string) *http.Request {
	data := url.Values{}

	data.Add("name", "seriessearch")
	data.Add("seriesId", seriesID)

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) RssSync() *http.Request {
	data := url.Values{}

	data.Add("name", "rsssync")

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) RenameFiles(files []int) *http.Request {
	data := url.Values{}

	data.Add("name", "renamefiles")

	if len(files) != 0 || files != nil {
		for _, id := range files {
			data.Add("files", string(id))
		}
	}

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) RenameSeries(seriesIds []int) *http.Request {
	data := url.Values{}

	data.Add("name", "renameseries")

	if len(seriesIds) != 0 || seriesIds != nil {
		for _, id := range seriesIds {
			data.Add("seriesIds", string(id))
		}
	}

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) DiskSpace() *http.Request {
	path := s.getAPIPath("diskspace", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) Backup() *http.Request {
	data := url.Values{}

	data.Add("name", "backup")

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) missingEpisodeSearch() *http.Request {
	data := url.Values{}

	data.Add("name", "missingEpisodeSearch")

	path := s.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) GetShow(seriesID string) *http.Request {
	data := url.Values{}

	data.Add("seriesId", seriesID)

	path := s.getAPIPath("episode", "")
	if req, err := http.NewRequest("GET", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) GetEpisode(episodeID string) *http.Request {
	path := s.getAPIPath("episode", episodeID)

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

// TODO Episode PUT to be implemented

func (s *Sonarr) GetEpisodeFiles(seriesID string) *http.Request {
	data := url.Values{}

	data.Add("seriesId", seriesID)

	path := s.getAPIPath("episodefile", "")

	if req, err := http.NewRequest("GET", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) GetShowFiles(id string) *http.Request {
	path := s.getAPIPath("episodefile", id)

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) DeleteEpisode(id string) *http.Request {
	path := s.getAPIPath("episodefile", id)

	if req, err := http.NewRequest("DELETE", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) History(sortKey, page, pageSize, sortDir string) *http.Request {
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

	path := s.getAPIPath("history", "")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) WantedMissing(sortKey, page, pageSize, sortDir string) *http.Request {
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

	path := s.getAPIPath("wanted", "missing")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) Queue() *http.Request {
	path := s.getAPIPath("queue", "")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) RemoveDownload(id string, blacklist bool) *http.Request {
	data := url.Values{}

	data.Add("id", id)

	if blacklist == true {
		data.Add("blacklist", "true")
	} else {
		data.Add("blacklist", "false")
	}

	path := s.getAPIPath("queue", "")

	if req, err := http.NewRequest("GET", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) GetProfiles() *http.Request {
	path := s.getAPIPath("profile", "")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}
