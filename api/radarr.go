package api

import (
	"bytes"
	"encoding/json"
	"jrs/config"
	"jrs/utils"
	"net/http"
	"net/url"
	"strings"
)

type image struct {
	coverType string
	url       string
}

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

func (r *Radarr) getAPIPath(args ...string) string {
	path := r.path + "/api"

	for _, arg := range args {
		path = path + "/" + arg
	}

	return path
}

func (r *Radarr) Calendar() *http.Request {
	// TODO start, end to be impletemented
	path := r.getAPIPath("calendar", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) DiskSpace() *http.Request {
	path := r.getAPIPath("diskspace", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) GetOngoingCommands() *http.Request {
	path := r.getAPIPath("command", "")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) GetCommandStatus(id string) *http.Request {
	path := r.getAPIPath("command", id)
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) RefreshMovie(movieID string) *http.Request {
	data := url.Values{}
	data.Add("name", "refreshmovie")

	if movieID != "" {
		data.Add("movieId", movieID)
	}

	path := r.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) RescanMovie(movieID string) *http.Request {
	data := url.Values{}
	data.Add("name", "rescanmovie")

	if movieID != "" {
		data.Add("movieId", movieID)
	}

	path := r.getAPIPath("command", "")
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

	path := r.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) DownloadedMoviesScan(path, downloadClient, importMode string) *http.Request {
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

	reqPath := r.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", reqPath, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil

}

func (r *Radarr) RssSync() *http.Request {
	data := url.Values{}

	data.Add("name", "rsssync")

	path := r.getAPIPath("command", "")
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

	path := r.getAPIPath("command", "")
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

	path := r.getAPIPath("command", "")
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

	path := r.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) NetImportSync() *http.Request {
	data := url.Values{}

	data.Add("name", "netimportsync")

	path := r.getAPIPath("command", "")
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

	path := r.getAPIPath("command", "")
	if req, err := http.NewRequest("POST", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) History(sortKey, page, pageSize, sortDir string) *http.Request {
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

	path := r.getAPIPath("history", "")

	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

// Nil id returns the all movies
func (r *Radarr) GetMovie(id string) *http.Request {
	path := r.getAPIPath("movie", id)
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

// Update and Add method not implemented yet || how to get rootfolder
func (r *Radarr) AddMovie(title, qualityProfileID, titleSlug, tmdbID, path string, images []image, monitored bool) *http.Request {
	data := url.Values{}
	img := new(bytes.Buffer)

	data.Add("title", title)
	data.Add("qualityProfileId", qualityProfileID)
	data.Add("titleSlug", titleSlug)
	data.Add("tmdbId ", tmdbID)

	if monitored == true {
		data.Add("monitored", "true")
	} else {
		data.Add("monitored", "false")
	}

	json.NewEncoder(img).Encode(images)
	data.Add("images", img.String())

	reqPath := r.getAPIPath("movie")
	if req, err := http.NewRequest("GET", reqPath, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) RemoveMovie(id string) *http.Request {
	path := r.getAPIPath("movie", id)
	if req, err := http.NewRequest("DELETE", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}

func (r *Radarr) SearchByName(name string) *http.Request {
	data := url.Values{}
	term := url.PathEscape(name)

	data.Add("term", term)

	path := r.getAPIPath("movie", "lookup")
	if req, err := http.NewRequest("GET", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil

}

func (r *Radarr) SearchByTmdb(tmdbID string) *http.Request {
	data := url.Values{}
	term := url.PathEscape(tmdbID)

	data.Add("tmdbId", term)

	path := r.getAPIPath("movie", "lookup", "tmdb")
	if req, err := http.NewRequest("GET", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil

}

func (r *Radarr) SearchByImdb(imdbID string) *http.Request {
	data := url.Values{}
	term := url.PathEscape(imdbID)

	data.Add("imdbId", term)

	path := r.getAPIPath("movie", "lookup", "imdb")
	if req, err := http.NewRequest("GET", path, strings.NewReader(data.Encode())); err == nil {
		req.Header = r.headers
		return req
	}
	return nil

}

func (r *Radarr) SystemStatus() *http.Request {
	path := r.getAPIPath("system", "status")
	if req, err := http.NewRequest("GET", path, nil); err == nil {
		req.Header = r.headers
		return req
	}
	return nil
}
