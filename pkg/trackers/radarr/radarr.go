package radarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"jrs/config"
	"jrs/pkg/trackers"
)

type Radarr struct {
	api     string
	path    string
	headers http.Header
}

func NewClient() *trackers.Client {
	c := new(trackers.Client)
	c.C = New(config.Params)
	c.Client = new(http.Client)
	return c
}

func New(c *config.Config) *Radarr {
	conf := c.GetDestination("Radarr")
	r := &Radarr{api: conf.Api, path: conf.Path, headers: http.Header{}}
	r.headers.Add("Content-Type", "application/json")
	r.headers.Add("X-Api-Key", r.api)
	return r
}

func (r *Radarr) BuildRequest(method string, body io.Reader, args ...string) (*http.Request, error) {
	path := r.path + "/api"

	for _, arg := range args {
		path = path + "/" + arg
	}

	if request, err := http.NewRequest(method, path, body); err == nil {
		request.Header = r.headers
		return request, err
	} else {
		return nil, err
	}
}

func (r *Radarr) Calendar() (*http.Request, error) {
	// TODO start, end to be implemented
	// path := r.getAPIPath("calendar", "")
	if req, err := r.BuildRequest("GET", nil, "calendar"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) DiskSpace() (*http.Request, error) {
	// path := r.getAPIPath("diskspace", "")
	if req, err := r.BuildRequest("GET", nil, "diskspace"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) GetOngoingCommands() (*http.Request, error) {
	// path := r.getAPIPath("command", "")
	if req, err := r.BuildRequest("GET", nil, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) GetCommandStatus(id string) (*http.Request, error) {
	// path := r.getAPIPath("command", id)
	if req, err := r.BuildRequest("GET", nil, "command", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RefreshMovie(movieID string) (*http.Request, error) {
	data := url.Values{}
	data.Add("name", "refreshmovie")

	if movieID != "" {
		data.Add("movieId", movieID)
	}

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) RescanMovie(movieID string) (*http.Request, error) {
	data := url.Values{}
	data.Add("name", "rescanmovie")

	if movieID != "" {
		data.Add("movieId", movieID)
	}

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) MoviesSearch(movieIds []int) (*http.Request, error) {
	data := url.Values{}
	data.Add("name", "moviessearch")

	if len(movieIds) != 0 || movieIds != nil {
		for _, id := range movieIds {
			data.Add("movieIds", string(id))
		}
	}

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) DownloadedMoviesScan(path, downloadClient, importMode string) (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "downloadedmoviesscan")

	if path != "" {
		data.Add("path", path)
	}

	if downloadClient != "" {
		// (nzoid for sabnzbd, special 'drone' attribute value for nzbget, uppercase infohash for torrents),
		data.Add("downloadClient", downloadClient)
	}

	if importMode != "" {
		data.Add("importMode", importMode)
	}

	// reqPath := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) RssSync() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "rsssync")

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) RenameFiles(files []int) (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "renamefiles")

	if len(files) != 0 || files != nil {
		for _, id := range files {
			data.Add("files", string(id))
		}
	}

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RenameMovies(movieIds []int) (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "renamemovies")

	if len(movieIds) != 0 || movieIds != nil {
		for _, id := range movieIds {
			data.Add("movieIds", string(id))
		}
	}

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) CutOffUnmetMoviesSearch(filterKey, filterValue string) (*http.Request, error) {
	data := url.Values{}

	if filterKey == "" {
		data.Add("filterKey", "monitored")
	} else {
		data.Add("filterKey", filterKey)
	}

	if filterValue == "" && filterKey == "" {
		data.Add("filterValue", "true") // (true (recommended), false), (all), (available, released, inCinemas, announced)
	} else {
		data.Add("filterValue", filterValue)
	}

	data.Add("name", "cutoffunmetmoviessearch") // monitored, all, status

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) NetImportSync() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "netimportsync")

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) MissingMoviesSearch(filterKey, filterValue string) (*http.Request, error) {
	data := url.Values{}

	if filterKey == "" {
		data.Add("filterKey", "monitored") // (Possible values: monitored (recommended), all, status)
	} else {
		data.Add("filterKey", filterKey)
	}

	if filterValue == "" && filterKey == "" {
		// (Possible values with respect to the ones for the filterKey above: (true (recommended), false), (all), (available, released, inCinemas, announced)
		data.Add("filterValue", "true")
	} else {
		data.Add("filterValue", filterValue)
	}

	data.Add("name", "missingmoviessearch")

	// path := r.getAPIPath("command", "")

	if req, err := r.BuildRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) History(sortKey, page, pageSize, sortDir string) (*http.Request, error) {
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

	if req, err := r.BuildRequest("GET", strings.NewReader(data.Encode()), "history"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// Nil id returns the all movies
func (r *Radarr) GetMovie(id string) (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "movie", id); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// Update and Add method not implemented yet || how to get rootfolder
func (r *Radarr) AddMovie(title, qualityProfileID, titleSlug, tmdbID, path string, images []trackers.Image, monitored bool) (*http.Request, error) {
	data := url.Values{}
	img := new(bytes.Buffer)

	data.Add("title", title)
	data.Add("qualityProfileId", qualityProfileID)
	data.Add("titleSlug", titleSlug)
	data.Add("tmdbId ", tmdbID)

	if monitored {
		data.Add("monitored", "true")
	} else {
		data.Add("monitored", "false")
	}

	err := json.NewEncoder(img).Encode(images)
	if err != nil {
		fmt.Printf("%v", err)
	}
	data.Add("images", img.String())

	// reqPath := r.getAPIPath("movie")

	if req, err := r.BuildRequest("GET", strings.NewReader(data.Encode()), "movie"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) RemoveMovie(id string) (*http.Request, error) {
	// path := r.getAPIPath("movie", id)

	if req, err := r.BuildRequest("DELETE", nil, "movie", id); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) SearchByName(name string) (*http.Request, error) {
	data := url.Values{}
	term := url.PathEscape(name)

	data.Add("term", term)

	// path := r.getAPIPath("movie", "lookup")

	if req, err := r.BuildRequest("GET", strings.NewReader(data.Encode()), "movie", "lookup"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) SearchByTmdb(tmdbID string) (*http.Request, error) {
	data := url.Values{}
	term := url.PathEscape(tmdbID)

	data.Add("tmdbId", term)

	// path := r.getAPIPath("movie", "lookup", "tmdb")

	if req, err := r.BuildRequest("GET", strings.NewReader(data.Encode()), "movie", "lookup", "tmdb"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) SearchByImdb(imdbID string) (*http.Request, error) {
	data := url.Values{}
	term := url.PathEscape(imdbID)

	data.Add("imdbId", term)

	// path := r.getAPIPath("movie", "lookup", "imdb")

	if req, err := r.BuildRequest("GET", strings.NewReader(data.Encode()), "movie", "lookup", "imdb"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) SystemStatus() (*http.Request, error) {
	// path := r.getAPIPath("system", "status")

	if req, err := r.BuildRequest("GET", nil, "system", "status"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) ConfigMediaManagement() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "config", "mediamanagement"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigNaming() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "config", "naming"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) ConfigIndexer() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "config", "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigDownloadClient() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "config", "downloadclient"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetDownloadClient() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "downloadclient"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigNotification() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "config", "notification"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigHost() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "config", "host"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigUI() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "config", "ui"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigNetImport() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "config", "netimport"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetNetImport() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "netimport"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigProfile() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "profile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigDelayProfile() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "delayprofile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) ConfigQualityDefinition() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "qualitydefinition"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetIndexers() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetRemotePathMapping() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "remotePathMapping"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetRootFolder() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "rootfolder"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetExclusions() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "exclusions"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetMetadata() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "metadata"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetIndexerSchema() (*http.Request, error) {
	if req, err := r.BuildRequest("GET", nil, "indexer", "schema"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) SetIndexer(i trackers.IndexerSchema) (*http.Request, error) {
	if req, err := r.BuildRequest("POST", nil, "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) DeleteIndexer(i trackers.IndexerSchema) (*http.Request, error) {
	id := fmt.Sprintf("%v", i.ID)
	if req, err := r.BuildRequest("DELETE", nil, "indexer", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) bulkImport(folderPath string, pageSize string) (*http.Request, error) {
	data := url.Values{}
	data.Add("id", "1")
	data.Add("folder", folderPath) // Full path to source folder
	data.Add("per_page", pageSize)
	data.Add("sort", "sortTitle")

	if req, err := r.BuildRequest("GET", strings.NewReader(data.Encode()), "movies", "bulkimport"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
