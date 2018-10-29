package radarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"jrs/config"
	"jrs/utils"
	"net/http"
	"net/url"
	"strings"
)

func New(c *config.Config) *Radarr {
	conf := c.GetDestination("Radarr")
	r := &Radarr{conf.API, utils.BuildURL(conf.IP, conf.Port), http.Header{}}
	r.headers.Add("Content-Type", "application/json")
	r.headers.Add("X-Api-Key", r.api)
	return r
}

func (r *Radarr) MakeRequest(method string, body io.Reader, args ...string) (*http.Request, error) {
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
	if req, err := r.MakeRequest("GET", nil, "calendar"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) DiskSpace() (*http.Request, error) {
	// path := r.getAPIPath("diskspace", "")
	if req, err := r.MakeRequest("GET", nil, "diskspace"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) GetOngoingCommands() (*http.Request, error) {
	// path := r.getAPIPath("command", "")
	if req, err := r.MakeRequest("GET", nil, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) GetCommandStatus(id string) (*http.Request, error) {
	// path := r.getAPIPath("command", id)
	if req, err := r.MakeRequest("GET", nil, "command", id); err == nil {
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

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
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

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
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

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
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
		data.Add("downloadClient", downloadClient) // (nzoid for sabnzbd, special 'drone' attribute value for nzbget, uppercase infohash for torrents),
	}

	if importMode != "" {
		data.Add("importMode", importMode)
	}

	// reqPath := r.getAPIPath("command", "")

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) RssSync() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "rsssync")

	// path := r.getAPIPath("command", "")

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
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

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
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

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
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

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) NetImportSync() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "netimportsync")

	// path := r.getAPIPath("command", "")

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
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
		data.Add("filterValue", "true") // (Possible values with respect to the ones for the filterKey above: (true (recommended), false), (all), (available, released, inCinemas, announced)
	} else {
		data.Add("filterValue", filterValue)
	}

	data.Add("name", "missingmoviessearch")

	// path := r.getAPIPath("command", "")

	if req, err := r.MakeRequest("POST", strings.NewReader(data.Encode()), "command"); err == nil {
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

	// path := r.getAPIPath("history", "")

	if req, err := r.MakeRequest("GET", strings.NewReader(data.Encode()), "history"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// Nil id returns the all movies
func (r *Radarr) GetMovie(id string) (*http.Request, error) {
	// path := r.getAPIPath("movie", id)

	if req, err := r.MakeRequest("GET", nil, "movie", id); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// Update and Add method not implemented yet || how to get rootfolder
func (r *Radarr) AddMovie(title, qualityProfileID, titleSlug, tmdbID, path string, images []image, monitored bool) (*http.Request, error) {
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

	// reqPath := r.getAPIPath("movie")

	if req, err := r.MakeRequest("GET", strings.NewReader(data.Encode()), "movie"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) RemoveMovie(id string) (*http.Request, error) {
	// path := r.getAPIPath("movie", id)

	if req, err := r.MakeRequest("DELETE", nil, "movie", id); err == nil {
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

	if req, err := r.MakeRequest("GET", strings.NewReader(data.Encode()), "movie", "lookup"); err == nil {
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

	if req, err := r.MakeRequest("GET", strings.NewReader(data.Encode()), "movie", "lookup", "tmdb"); err == nil {
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

	if req, err := r.MakeRequest("GET", strings.NewReader(data.Encode()), "movie", "lookup", "imdb"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) SystemStatus() (*http.Request, error) {
	// path := r.getAPIPath("system", "status")

	if req, err := r.MakeRequest("GET", nil, "system", "status"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// GET /config/mediamanagement Response
// {
// "autoUnmonitorPreviouslyDownloadedEpisodes": false,
// "recycleBin": "",
// "autoDownloadPropers": true,
// "createEmptySeriesFolders": false,
// "fileDate": "none",
// "autoRenameFolders": false,
// "pathsDefaultStatic": false,
// "setPermissionsLinux": false,
// "fileChmod": "0644",
// "folderChmod": "0755",
// "chownUser": "",
// "chownGroup": "",
// "skipFreeSpaceCheckWhenImporting": false,
// "copyUsingHardlinks": false,
// "importExtraFiles": false,
// "extraFileExtensions": "srt",
// "enableMediaInfo": true,
// "id": 1
// }
func (r *Radarr) ConfigMediaManagement() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "config", "mediamanagement"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /config/naming Response
// {
// "renameEpisodes": true,
// "replaceIllegalCharacters": true,
// "colonReplacementFormat": "delete",
// "standardMovieFormat": "{Movie.Title}.{Release.Year}.{Quality.Title}",
// "movieFolderFormat": "{Movie Title} ({Release Year})",
// "multiEpisodeStyle": 0,
// "includeSeriesTitle": false,
// "includeEpisodeTitle": false,
// "includeQuality": false,
// "replaceSpaces": false,
// "id": 1
// }
func (r *Radarr) ConfigNaming() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "config", "naming"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// GET /config/indexer
// {
// "minimumAge": 0,
// "maximumSize": 0,
// "retention": 0,
// "rssSyncInterval": 60,
// "preferIndexerFlags": true,
// "availabilityDelay": 0,
// "allowHardcodedSubs": false,
// "whitelistedHardcodedSubs": "",
// "parsingLeniency": "strict",
// "id": 1
// }
func (r *Radarr) ConfigIndexer() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "config", "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /config/downloadclient
// {
// "downloadedMoviesFolder": "",
// "downloadClientWorkingFolders": "_UNPACK_|_FAILED_",
// "downloadedMoviesScanInterval": 0,
// "enableCompletedDownloadHandling": true,
// "removeCompletedDownloads": true,
// "autoRedownloadFailed": true,
// "removeFailedDownloads": true,
// "id": 1
// }
func (r *Radarr) ConfigDownloadClient() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "config", "downloadclient"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /downloadclient
// [{
// "enable": true,
// "protocol": "torrent",
// "name": "transmission",
// "fields": [
// {
// "order": 0,
// "name": "Host",
// "label": "Host",
// "value": "192.*.*.*",
// "type": "textbox",
// "advanced": false
// },
// {
// "order": 1,
// "name": "Port",
// "label": "Port",
// "value": 9091,
// "type": "textbox",
// "advanced": false
// },
// {
// "order": 2,
// "name": "UrlBase",
// "label": "Url Base",
// "helpText": "Adds a prefix to the transmission rpc url, eg http://[host]:[port]/[urlBase]/rpc, defaults to '/transmission/'",
// "value": "/transmission/",
// "type": "textbox",
// "advanced": true
// },
// {
// "order": 3,
// "name": "Username",
// "label": "Username",
// "type": "textbox",
// "advanced": false
// },
// {
// "order": 4,
// "name": "Password",
// "label": "Password",
// "type": "password",
// "advanced": false
// },
// {
// "order": 5,
// "name": "TvCategory",
// "label": "Category",
// "helpText": "Adding a category specific to Sonarr avoids conflicts with unrelated downloads, but it's optional. Creates a [category] subdirectory in the output directory.",
// "type": "textbox",
// "advanced": false
// },
// {
// "order": 6,
// "name": "TvDirectory",
// "label": "Directory",
// "helpText": "Optional location to put downloads in, leave blank to use the default Transmission location",
// "type": "textbox",
// "advanced": true
// },
// {
// "order": 7,
// "name": "RecentTvPriority",
// "label": "Recent Priority",
// "helpText": "Priority to use when grabbing episodes that aired within the last 14 days",
// "value": 0,
// "type": "select",
// "advanced": false,
// "selectOptions": [
// {
// "value": 0,
// "name": "Last"
// },
// {
// "value": 1,
// "name": "First"
// }
// ]},
// {
// "order": 8,
// "name": "OlderTvPriority",
// "label": "Older Priority",
// "helpText": "Priority to use when grabbing episodes that aired over 14 days ago",
// "value": 0,
// "type": "select",
// "advanced": false,
// "selectOptions": [
// {
// "value": 0,
// "name": "Last"
// },
// {
// "value": 1,
// "name": "First"
// }]
// },
// {
// "order": 9,
// "name": "AddPaused",
// "label": "Add Paused",
// "value": false,
// "type": "checkbox",
// "advanced": false
// },
// {
// "order": 10,
// "name": "UseSsl",
// "label": "Use SSL",
// "value": false,
// "type": "checkbox",
// "advanced": false
// }],
// "implementationName": "Transmission",
// "implementation": "Transmission",
// "configContract": "TransmissionSettings",
// "infoLink": "https://github.com/Sonarr/Sonarr/wiki/Supported-DownloadClients#transmission",
// "id": 1
// }]
func (r *Radarr) GetDownloadClient() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "downloadclient"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /config/notification
// {
// "downloadedMoviesFolder": "",
// "downloadClientWorkingFolders": "_UNPACK_|_FAILED_",
// "downloadedMoviesScanInterval": 0,
// "enableCompletedDownloadHandling": true,
// "removeCompletedDownloads": true,
// "autoRedownloadFailed": true,
// "removeFailedDownloads": true,
// "id": 1
// }
func (r *Radarr) ConfigNotification() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "config", "notification"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /config/host
// {
// "bindAddress": "*",
// "port": 9090,
// "sslPort": 9898,
// "enableSsl": false,
// "launchBrowser": true,
// "authenticationMethod": "none",
// "analyticsEnabled": true,
// "logLevel": "Info",
// "branch": "nightly",
// "apiKey": "XXXXXXXXXXXXX",
// "sslCertHash": "",
// "urlBase": "/radarr",
// "updateAutomatically": false,
// "updateMechanism": "builtIn",
// "updateScriptPath": "",
// "proxyEnabled": false,
// "proxyType": "http",
// "proxyHostname": "",
// "proxyPort": 8080,
// "proxyUsername": "",
// "proxyPassword": "",
// "proxyBypassFilter": "",
// "proxyBypassLocalAddresses": true,
// "id": 1
// }
func (r *Radarr) ConfigHost() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "config", "host"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /config/ui
// {
// "firstDayOfWeek": 1,
// "calendarWeekColumnHeader": "ddd M/D",
// "shortDateFormat": "MMM D YYYY",
// "longDateFormat": "dddd, MMMM D YYYY",
// "timeFormat": "h(:mm)a",
// "showRelativeDates": true,
// "enableColorImpairedMode": false,
// "id": 1
// }
func (r *Radarr) ConfigUI() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "config", "ui"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /config/netimport
// {
// "netImportSyncInterval": 1440,
// "listSyncLevel": "disabled",
// "importExclusions": "",
// "traktAuthToken": "XXX",
// "traktRefreshToken": "XXX",
// "traktTokenExpiry": 1532979325,
// "id": 1
// }
func (r *Radarr) ConfigNetImport() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "config", "netimport"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetNetImport() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "netimport"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /config/profile
// Array of below
// {
// "name": "Any",
// "cutoff": {
// "id": 4,
// "name": "HDTV-720p",
// "source": "television",
// "resolution": 720
// },
// "items": [
// {
// "quality": {
// "id": 0,
// "name": "Unknown",
// "source": "unknown",
// "resolution": 0
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 1,
// "name": "SDTV",
// "source": "television",
// "resolution": 480
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 8,
// "name": "WEBDL-480p",
// "source": "web",
// "resolution": 480
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 2,
// "name": "DVD",
// "source": "dvd",
// "resolution": 480
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 4,
// "name": "HDTV-720p",
// "source": "television",
// "resolution": 720
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 9,
// "name": "HDTV-1080p",
// "source": "television",
// "resolution": 1080
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 10,
// "name": "Raw-HD",
// "source": "televisionRaw",
// "resolution": 1080
// },
// "allowed": false
// },
// {
// "quality": {
// "id": 5,
// "name": "WEBDL-720p",
// "source": "web",
// "resolution": 720
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 6,
// "name": "Bluray-720p",
// "source": "bluray",
// "resolution": 720
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 3,
// "name": "WEBDL-1080p",
// "source": "web",
// "resolution": 1080
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 7,
// "name": "Bluray-1080p",
// "source": "bluray",
// "resolution": 1080
// },
// "allowed": true
// },
// {
// "quality": {
// "id": 16,
// "name": "HDTV-2160p",
// "source": "television",
// "resolution": 2160
// },
// "allowed": false
// },
// {
// "quality": {
// "id": 18,
// "name": "WEBDL-2160p",
// "source": "web",
// "resolution": 2160
// },
// "allowed": false
// },
// {
// "quality": {
// "id": 19,
// "name": "Bluray-2160p",
// "source": "bluray",
// "resolution": 2160
// },
// "allowed": false
// }
// ],
// "language": "english",
// "id": 1
// }
func (r *Radarr) ConfigProfile() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "profile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /delayprofile
// [{
// "enableUsenet": false,
// "enableTorrent": true,
// "preferredProtocol": "torrent",
// "usenetDelay": 0,
// "torrentDelay": 120,
// "order": 2147483647,
// "tags": [],
// "id": 1
// }
// ]
func (r *Radarr) ConfigDelayProfile() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "delayprofile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /qualitydefinition
// [{
// "quality": {
// "id": 0,
// "name": "Unknown",
// "source": "unknown",
// "resolution": 0
// },
// "title": "Unknown",
// "weight": 1,
// "minSize": 0.0,
// "maxSize": 40.9,
// "id": 1
// },
// {
// "quality": {
// "id": 1,
// "name": "SDTV",
// "source": "television",
// "resolution": 480
// },
// "title": "SDTV",
// "weight": 2,
// "minSize": 0.0,
// "maxSize": 11.9,
// "id": 2
// },
// {
// "quality": {
// "id": 8,
// "name": "WEBDL-480p",
// "source": "web",
// "resolution": 480
// },
// "title": "WEBDL-480p",
// "weight": 3,
// "minSize": 0.0,
// "maxSize": 19.1,
// "id": 3
// },
// {
// "quality": {
// "id": 2,
// "name": "DVD",
// "source": "dvd",
// "resolution": 480
// },
// "title": "DVD",
// "weight": 4,
// "minSize": 0.0,
// "maxSize": 13.9,
// "id": 4
// },
// {
// "quality": {
// "id": 4,
// "name": "HDTV-720p",
// "source": "television",
// "resolution": 720
// },
// "title": "HDTV-720p",
// "weight": 5,
// "minSize": 0.0,
// "maxSize": 27.5,
// "id": 5
// },
// {
// "quality": {
// "id": 9,
// "name": "HDTV-1080p",
// "source": "television",
// "resolution": 1080
// },
// "title": "HDTV-1080p",
// "weight": 6,
// "minSize": 0.0,
// "maxSize": 44.3,
// "id": 6
// },
// {
// "quality": {
// "id": 10,
// "name": "Raw-HD",
// "source": "televisionRaw",
// "resolution": 1080
// },
// "title": "Raw-HD",
// "weight": 7,
// "minSize": 0.0,
// "id": 7
// },
// {
// "quality": {
// "id": 5,
// "name": "WEBDL-720p",
// "source": "web",
// "resolution": 720
// },
// "title": "WEBDL-720p",
// "weight": 8,
// "minSize": 0.0,
// "maxSize": 30.1,
// "id": 8
// },
// {
// "quality": {
// "id": 6,
// "name": "Bluray-720p",
// "source": "bluray",
// "resolution": 720
// },
// "title": "Bluray-720p",
// "weight": 9,
// "minSize": 0.0,
// "maxSize": 28.8,
// "id": 9
// },
// {
// "quality": {
// "id": 3,
// "name": "WEBDL-1080p",
// "source": "web",
// "resolution": 1080
// },
// "title": "WEBDL-1080p",
// "weight": 10,
// "minSize": 0.0,
// "maxSize": 44.3,
// "id": 10
// },
// {
// "quality": {
// "id": 7,
// "name": "Bluray-1080p",
// "source": "bluray",
// "resolution": 1080
// },
// "title": "Bluray-1080p",
// "weight": 11,
// "minSize": 0.0,
// "maxSize": 44.7,
// "id": 11
// },
// {
// "quality": {
// "id": 16,
// "name": "HDTV-2160p",
// "source": "television",
// "resolution": 2160
// },
// "title": "HDTV-2160p",
// "weight": 12,
// "minSize": 0.0,
// "id": 12
// },
// {
// "quality": {
// "id": 18,
// "name": "WEBDL-2160p",
// "source": "web",
// "resolution": 2160
// },
// "title": "WEBDL-2160p",
// "weight": 13,
// "minSize": 0.0,
// "id": 13
// },
// {
// "quality": {
// "id": 19,
// "name": "Bluray-2160p",
// "source": "bluray",
// "resolution": 2160
// },
// "title": "Bluray-2160p",
// "weight": 14,
// "minSize": 0.0,
// "id": 14
// }
// ]
func (r *Radarr) ConfigQualityDefinition() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "qualitydefinition"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /indexers
// {
// "enableRss": true,
// "enableSearch": true,
// "supportsRss": true,
// "supportsSearch": true,
// "protocol": "torrent",
// "name": "demon",
// "fields": [
// {
// "order": 0,
// "name": "BaseUrl",
// "label": "URL",
// "value": "http://192.168.1.**:8888/torznab/demonoid",
// "type": "textbox",
// "advanced": false
// },
// {
// "order": 1,
// "name": "ApiPath",
// "label": "API Path",
// "helpText": "Path to the api, usually /api",
// "value": "/api",
// "type": "textbox",
// "advanced": true
// },
// {
// "order": 2,
// "name": "ApiKey",
// "label": "API Key",
// "value": "API KEY",
// "type": "textbox",
// "advanced": false
// },
// {
// "order": 3,
// "name": "Categories",
// "label": "Categories",
// "helpText": "Comma Separated list, leave blank to disable standard/daily shows",
// "value": [
// 5030,
// 5040
// ],
// "type": "textbox",
// "advanced": false
// },
// {
// "order": 4,
// "name": "AnimeCategories",
// "label": "Anime Categories",
// "helpText": "Comma Separated list, leave blank to disable anime",
// "value": [],
// "type": "textbox",
// "advanced": false
// },
// {
// "order": 5,
// "name": "AdditionalParameters",
// "label": "Additional Parameters",
// "helpText": "Additional Newznab parameters",
// "type": "textbox",
// "advanced": true
// },
// {
// "order": 6,
// "name": "MinimumSeeders",
// "label": "Minimum Seeders",
// "helpText": "Minimum number of seeders required.",
// "value": 1,
// "type": "textbox",
// "advanced": true
// },
// {
// "order": 7,
// "name": "SeedCriteria.SeedRatio",
// "label": "Seed Ratio",
// "helpText": "The ratio a torrent should reach before stopping, empty is download client's default",
// "type": "textbox",
// "advanced": true
// },
// {
// "order": 8,
// "name": "SeedCriteria.SeedTime",
// "label": "Seed Time",
// "unit": "minutes",
// "helpText": "The time a torrent should be seeded before stopping, empty is download client's default",
// "type": "textbox",
// "advanced": true
// },
// {
// "order": 9,
// "name": "SeedCriteria.SeasonPackSeedTime",
// "label": "Season-Pack Seed Time",
// "unit": "minutes",
// "helpText": "The time a torrent should be seeded before stopping, empty is download client's default",
// "type": "textbox",
// "advanced": true
// }
// ],
// "implementationName": "Torznab",
// "implementation": "Torznab",
// "configContract": "TorznabSettings",
// "infoLink": "https://github.com/Sonarr/Sonarr/wiki/Supported-Indexers#torznab",
// "id": 13
// }
func (r *Radarr) GetIndexers() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// GET /remotePathMapping
// [{
// "host": "192.*.*.*",
// "remotePath": "/*/*/*/*/",
// "localPath": "/*/*/",
// "id": 1
// }]
func (r *Radarr) GetRemotePathMapping() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "remotePathMapping"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetRootFolder() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "rootfolder"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetExclusions() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "exclusions"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetMetadata() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "metadata"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetIndexerSchema() (*http.Request, error) {
	if req, err := r.MakeRequest("GET", nil, "indexer", "schema"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) SetIndexer(i IndexerSchema) (*http.Request, error) {
	if req, err := r.MakeRequest("POST", nil, "indexer"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) DeleteIndexer(i IndexerSchema) (*http.Request, error) {
	id := fmt.Sprintf("%v", i.ID)
	if req, err := r.MakeRequest("DELETE", nil, "indexer", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
