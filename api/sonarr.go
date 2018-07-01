package api

import (
	"jrs/config"
	"net/http"
)

type Sonarr struct {
	api     string
	path    string
	headers http.Header
}

func NewSonarr(conf *config.Config) *Sonarr {
	s := new(Sonarr)
	s.headers.Add("Content-Type", "application/json")
	// s.headers.Add("X-Api-Key")
	return s

}

func (s *Sonarr) getApiPath(endpoint, action string, args ...string) string {
	path := s.path + "/api/" + endpoint + "?apikey=" + s.api

	return path
}

func (s *Sonarr) Calendar() *http.Request {
	path := s.getApiPath("calendar", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) DiskSpace() *http.Request {
	path := s.getApiPath("diskspace", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) GetSeries() *http.Request {
	path := s.getApiPath("series", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}

func (s *Sonarr) GetSerie(id string) *http.Request {
	path := s.getApiPath("series", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = s.headers
		return req
	}
	return nil
}
