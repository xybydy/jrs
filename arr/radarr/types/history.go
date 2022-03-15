package types

import "time"

type History struct {
	Page          int             `json:"page"`
	PageSize      int             `json:"pageSize"`
	SortKey       string          `json:"sortKey"`
	SortDirection string          `json:"sortDirection"`
	TotalRecords  int             `json:"totalRecords"`
	Records       []HistoryRecord `json:"records"`
}

type HistoryRecord struct {
	ID                  int64          `json:"id"`
	MovieID             int64          `json:"movieId"`
	SourceTitle         string         `json:"sourceTitle"`
	Languages           []Language     `json:"languages"`
	Quality             Quality        `json:"quality"`
	CustomFormats       []CustomFormat `json:"customFormats"`
	QualityCutoffNotMet bool           `json:"qualityCutoffNotMet"`
	Date                time.Time      `json:"date"`
	DownloadID          string         `json:"downloadId"`
	EventType           string         `json:"eventType"`
	Data                struct {
		Age                string    `json:"age"`
		AgeHours           string    `json:"ageHours"`
		AgeMinutes         string    `json:"ageMinutes"`
		DownloadClient     string    `json:"downloadClient"`
		DownloadClientName string    `json:"downloadClientName"`
		DownloadURL        string    `json:"downloadUrl"`
		DroppedPath        string    `json:"droppedPath"`
		FileID             string    `json:"fileId"`
		GUID               string    `json:"guid"`
		ImportedPath       string    `json:"importedPath"`
		Indexer            string    `json:"indexer"`
		IndexerFlags       string    `json:"indexerFlags"`
		IndexerID          string    `json:"indexerId"`
		Message            string    `json:"message"`
		NzbInfoURL         string    `json:"nzbInfoUrl"`
		Protocol           string    `json:"protocol"`
		PublishedDate      time.Time `json:"publishedDate"`
		Reason             string    `json:"reason"`
		ReleaseGroup       string    `json:"releaseGroup"`
		Size               string    `json:"size"`
		TmdbID             string    `json:"tmdbId"`
		TorrentInfoHash    string    `json:"torrentInfoHash"`
	} `json:"data"`
}
