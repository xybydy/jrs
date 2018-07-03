package api

import (
	"jrs/config"
	"jrs/utils"
	"net/http"
	"net/url"
	"strings"
)

type Radarr struct {
	api     string
	path    string
	headers http.Header
}

func NewRadarr(c *config.Config) *Radarr {
	conf := c.GetDestination("Radarr")
	r := &Radarr{conf.Api, utils.BuildURL(conf.Ip, conf.Port), http.Header{}}
	r.headers.Add("Content-Type", "application/json")
	r.headers.Add("X-Api-Key", r.api)
	return r
}

func (r *Radarr) getApiPath(endpoint, action string, args ...string) string {
	path := r.path + "/api/" + endpoint

	if action != "" {
		path = path + "/" + action
	}

	return path
}

func (r *Radarr) Calendar() *http.Request {
	// TODO start, end to be impletemented
	path := r.getApiPath("calendar", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) DiskSpace() *http.Request {
	path := r.getApiPath("diskspace", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) GetOngoingCommands() *http.Request {
	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) GetCommandStatus(id string) *http.Request {
	path := r.getApiPath("command", id)
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) RefreshMovie(movieId string) *http.Request {
	data := url.Values{}
	data.Add("name", "refreshmovie")

	if movieId != "" {
		data.Add("movieId", movieId)
	}

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) RescanMovie(movieId string) *http.Request {
	data := url.Values{}
	data.Add("name", "rescanmovie")

	if movieId != "" {
		data.Add("movieId", movieId)
	}

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) MoviesSearch(movieIds []int) *http.Request {
	data := url.Values{}
	data.Add("name", "moviessearch")

	if len(movieIds) != 0 || movieIds != nil {
		for _, id := range movieIds {
			data.Add("movieIds", string(id))
		}
	}

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) DownloadedMoviesScan(path, downloadClient, importMode string) {
	data := url.Values{}

	data.Add("name", "downloadedmoviesscan")

	if path != "" {
		data.Add("path", path)
	}

	if downloadClient != "" {
		data.Add("downloadClient", downloadClient) //(nzoid for sabnzbd, special 'drone' attribute value for nzbget, uppercase infohash for torrents),
	}

	if importMode != "" {
		data.Add("importMode", importMode)
	}

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil

}

func (r *Radarr) RssSync() *http.Request {
	data := url.Values{}

	data.Add("name", "rsssync")

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) RenameFiles(files []int) *http.Request {
	data := url.Values{}

	data.Add("name", "renamefiles")

	if len(files) != 0 || files != nil {
		for _, id := range files {
			data.Add("files", string(id))
		}
	}

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) RenameMovies(movieIds []int) *http.Request {
	data := url.Values{}

	data.Add("name", "renamemovies")

	if len(movieIds) != 0 || movieIds != nil {
		for _, id := range movieIds {
			data.Add("movieIds", string(id))
		}
	}

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) CutOffUnmetMoviesSearch(filterKey, filterValue string) *http.Request {
	data := url.Values{}

	if filterKey == "" {
		data.Add("filterKey", "monitored")
	} else {
		data.Add("filterKey", filterKey)
	}

	if filterValue == "" && filterKey == "" {
		data.Add("filterValue", "true") //(true (recommended), false), (all), (available, released, inCinemas, announced)
	} else {
		data.Add("filterValue", filterValue)
	}

	data.Add("name", "cutoffunmetmoviessearch") //monitored, all, status

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) NetImportSync() *http.Request {
	data := url.Values{}

	data.Add("name", "netimportsync")

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) MissingMoviesSearch(filterKey, filterValue string) *http.Request {
	data := url.Values{}

	if filterKey == "" {
		data.Add("filterKey", "monitored") // (Possible values: monitored (recommended), all, status)
	} else {
		data.Add("filterKey", filterKey)
	}

	if filterValue == "" && filterKey == "" {
		data.Add("filterValue", "true") //(Possible values with respect to the ones for the filterKey above: (true (recommended), false), (all), (available, released, inCinemas, announced)
	} else {
		data.Add("filterValue", filterValue)
	}

	data.Add("name", "missingmoviessearch")

	path := r.getApiPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}
