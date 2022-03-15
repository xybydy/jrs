package radarr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/xybydy/jrs/arr"
	"github.com/xybydy/jrs/arr/radarr/types"
	"github.com/xybydy/jrs/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const apiVer = "v3"

type Radarr struct {
	*arr.Arr
}

func New(url, apiKey string) *Radarr {
	r := new(Radarr)
	r.URL = url
	r.ApiKey = apiKey
	r.ApiVer = apiVer
	r.Timeout = 60
	return r
}

// MOVIE

// Nil id returns the all movies
func (r *Radarr) GetMovie(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "movie", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Update and Add method not implemented yet || how to get rootfolder
func (r *Radarr) AddMovie(movie types.Movie) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(movie); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "movie"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) EditMovie(movie types.Movie, moveFiles bool) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(movie); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPut, &body, "movie"); err == nil {
		query := req.URL.Query()
		query.Add("moveFiles", strconv.FormatBool(moveFiles))
		req.URL.RawQuery = query.Encode()
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) RemoveMovie(id string, addImportExclusion bool, deleteFiles bool) (*http.Request, error) {

	if req, err := r.Request(http.MethodDelete, nil, "movie", id); err == nil {
		query := req.URL.Query()
		query.Add("addImportExclusion", strconv.FormatBool(addImportExclusion))
		query.Add("deleteFiles", strconv.FormatBool(deleteFiles))
		req.URL.RawQuery = query.Encode()
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) LookupMovie(name string) (*http.Request, error) {
	data := url.Values{}
	term := url.PathEscape(name)

	data.Add("term", term)

	if req, err := r.Request(http.MethodGet, strings.NewReader(data.Encode()), "movie", "lookup"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) EditMultipleMovies(param types.MovieEditorEditBody) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(param); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodPut, &body, "movie", "editor"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) RemoveMultipleMovies(param types.MovieEditorRemoveBody) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(param); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodDelete, &body, "movie", "editor"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) AddMultipleMovies(movies ...types.Movie) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(movies); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodPost, &body, "movie", "import"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// MovieFile

// Nil id returns the all
func (r *Radarr) GetMovieFile(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "moviefile", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RemoveMovieFile(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodDelete, nil, "moviefile", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// History
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

	if req, err := r.Request(http.MethodGet, strings.NewReader(data.Encode()), "history"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) HistoryMovie(id, eventType int64) (*http.Request, error) {
	data := url.Values{}
	data.Add("movieId", strconv.FormatInt(id, 10))
	data.Add("eventType", strconv.FormatInt(eventType, 10))

	if req, err := r.Request(http.MethodGet, strings.NewReader(data.Encode()), "history", "movie"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Blocklist

func (r *Radarr) GetBlocklist(sortKey, page, pageSize, sortDir string) (*http.Request, error) {
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

	if req, err := r.Request(http.MethodGet, strings.NewReader(data.Encode()), "blocklist"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RemoveBlocklist(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodDelete, nil, "blocklist", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetBlocklistMovie(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "blocklist", "movie", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RemoveBulkBlocklist() (*http.Request, error) {
	if req, err := r.Request(http.MethodDelete, nil, "blocklist", "bulk"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Qeueu
func (r *Radarr) GetQueue(sortKey, page, pageSize, sortDir string, includeUnknownMovieItems bool) (*http.Request, error) {
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

	data.Add("includeUnknownMovieItems", strconv.FormatBool(includeUnknownMovieItems))

	if req, err := r.Request(http.MethodGet, strings.NewReader(data.Encode()), "queue"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RemoveQueue(id string, removeFromClient, blocklist bool) (*http.Request, error) {
	data := url.Values{}
	data.Add("removeFromClient", strconv.FormatBool(removeFromClient))
	data.Add("blocklist", strconv.FormatBool(blocklist))

	if req, err := r.Request(http.MethodDelete, strings.NewReader(data.Encode()), "queue", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RemoveBulkQueue(removeFromClient, blocklist bool) (*http.Request, error) {
	data := url.Values{}
	data.Add("removeFromClient", strconv.FormatBool(removeFromClient))
	data.Add("blocklist", strconv.FormatBool(blocklist))

	if req, err := r.Request(http.MethodDelete, strings.NewReader(data.Encode()), "queue", "bulk"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) QueueDetails(includeMovie bool) (*http.Request, error) {
	data := url.Values{}
	data.Add("includeMovie", strconv.FormatBool(includeMovie))

	if req, err := r.Request(http.MethodGet, strings.NewReader(data.Encode()), "queue", "details"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) QueueStatus() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "queue", "status"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) QueueGrab(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodPost, nil, "queue", "grab", "id"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Indexer

// nil gives all
func (r *Radarr) GetIndexer(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "indexer", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) EditIndexer(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodPut, nil, "indexer", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RemoveIndexer(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodDelete, nil, "indexer", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetIndexerSchema() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "indexer", "schema"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Download Client

// Nil id gives all
func (r *Radarr) GetDownloadClient(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "downloadclient", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RemoveDownloadClient(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodDelete, nil, "downloadclient", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) EditDownloadClient(client types.DownloadClient) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(client); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPut, &body, "downloadclient", fmt.Sprintf("%d", client.Id)); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// ImportList

// Nil id gives all
func (r *Radarr) GetImportList(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "importlist", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) RemoveImportList(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodDelete, nil, "importlist", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) EditImportList(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodPut, nil, "importlist", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Notification

func (r *Radarr) GetNotification(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "notification", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) RemoveNotification(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodDelete, nil, "notification", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) EditNotification(n types.Notification) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(n); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodPut, &body, "notification", fmt.Sprintf("%d'", n.Id)); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Tag

func (r *Radarr) GetTagDetails(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "tag", "detail", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetTag(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "tag", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) RemoveTag(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodDelete, nil, "tag", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) EditTag(t types.Tag) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(t); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodPut, &body, "tag", fmt.Sprintf("%d", t.ID)); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) CreateTag(t types.Tag) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(t); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodPost, &body, "tag"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) DiskSpace() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "diskspace"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Config

func (r *Radarr) GetConfigUI() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "config", "ui"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) EditConfigUI(c types.ConfigUI) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(c); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodPut, &body, "config", "ui"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetConfigHost() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "config", "host"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) EditConfigHost(c types.ConfigHost) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(c); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodPut, &body, "config", "host"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetConfigNaming() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "config", "naming"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) EditConfigNaming(c types.ConfigNaming) (*http.Request, error) {
	body := bytes.Buffer{}

	if err := json.NewEncoder(&body).Encode(c); err != nil {
		return nil, err
	}
	if req, err := r.Request(http.MethodPut, &body, "config", "host"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetMetadata() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "metadata"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetSystemStatus() (*http.Request, error) {
	// path := r.getAPIPath("system", "status")

	if req, err := r.Request(http.MethodGet, nil, "system", "status"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

func (r *Radarr) GetHealth() (*http.Request, error) {
	// path := r.getAPIPath("system", "status")

	if req, err := r.Request(http.MethodGet, nil, "health"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}

// Commands

func (r *Radarr) CommandApplicationUpdate() (*http.Request, error) {
	req := struct{ name string }{"ApplicationUpdate"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}

}
func (r *Radarr) CommandBackup() (*http.Request, error) {
	req := struct{ name string }{"Backup"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandClearBlocklist() (*http.Request, error) {
	req := struct{ name string }{"ClearBlocklist"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandCleanUpRecycleBin() (*http.Request, error) {
	req := struct{ name string }{"CleanUpRecycleBin"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandCutoffUnmetMoviesSearch() (*http.Request, error) {
	req := struct{ name string }{"CutoffUnmetMoviesSearch"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandDeleteLogFiles() (*http.Request, error) {
	req := struct{ name string }{"DeleteLogFiles"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandDeleteUpdateLogFiles() (*http.Request, error) {
	req := struct{ name string }{"DeleteUpdateLogFiles"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandDownloadedMoviesScan() (*http.Request, error) {
	req := struct{ name string }{"DownloadedMoviesScan"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandMissingMoviesSearch() (*http.Request, error) {
	req := struct{ name string }{"MissingMoviesSearch"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandManualImport() (*http.Request, error) {
	req := struct{ name string }{"ManualImport"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandRefreshMonitoredDownloads() (*http.Request, error) {
	req := struct{ name string }{"RefreshMonitoredDownloads"}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandRefreshMovie(ids ...string) (*http.Request, error) {
	req := struct {
		name     string
		movieIds []string
	}{}

	req.name = "RefreshMovie"

	for _, id := range ids {
		req.movieIds = append(req.movieIds, id)
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandRenameFiles(id string) (*http.Request, error) {
	req := struct {
		name    string
		movieId string
	}{
		"RenameFiles",
		id,
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandRenameMovie() (*http.Request, error) {
	req := struct {
		name string
	}{
		"RenameMovie",
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandResetApiKey() (*http.Request, error) {
	req := struct {
		name string
	}{
		"ResetApiKey",
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandMoviesSearch(ids ...string) (*http.Request, error) {
	req := struct {
		name     string
		movieIds []string
	}{}

	req.name = "MoviesSearch"

	for _, id := range ids {
		req.movieIds = append(req.movieIds, id)
	}

	body := bytes.Buffer{}
	if err := json.NewEncoder(&body).Encode(req); err != nil {
		return nil, err
	}

	if req, err := r.Request(http.MethodPost, &body, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandRssSync() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "rsssync")

	if req, err := r.Request(http.MethodPost, strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) CommandImportListSync() (*http.Request, error) {
	data := url.Values{}

	data.Add("name", "ImportListSync")

	if req, err := r.Request(http.MethodPost, strings.NewReader(data.Encode()), "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetOngoingCommands() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "command"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
func (r *Radarr) GetCommandStatus(id string) (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "command", id); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetUpdate() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "update"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetQualityProfile() (*http.Request, error) {
	if req, err := r.Request(http.MethodGet, nil, "qualityprofile"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

// Dateformat is YYYY-MM-DD
func (r *Radarr) GetCalendar(unmonitored bool, start, end string) (*http.Request, error) {
	data := url.Values{}
	data.Add("unmonitored", strconv.FormatBool(unmonitored))

	if start == "" {
		start = time.Now().Format(utils.Iso86001)
	}

	if end == "" {
		end = time.Now().AddDate(0, 1, 0).Format(utils.Iso86001)
	}

	data.Add("start", start)
	data.Add("end", end)

	if req, err := r.Request(http.MethodGet, nil, "calendar"); err == nil {
		req.URL.RawQuery = data.Encode()
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetCustomFilter() (*http.Request, error) {

	if req, err := r.Request(http.MethodGet, nil, "customfilter"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetRemoteMapping() (*http.Request, error) {

	if req, err := r.Request(http.MethodGet, nil, "remotemapping"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}

func (r *Radarr) GetRootFolder() (*http.Request, error) {

	if req, err := r.Request(http.MethodGet, nil, "rootfolder"); err == nil {
		return req, err
	} else {
		return nil, err
	}
}
