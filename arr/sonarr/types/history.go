package types

import "time"

type History struct {
	Page          int    `json:"page"`
	PageSize      int    `json:"pageSize"`
	SortKey       string `json:"sortKey"`
	SortDirection string `json:"sortDirection"`
	TotalRecords  int    `json:"totalRecords"`
	Records       []struct {
		EpisodeId   int    `json:"episodeId"`
		SeriesId    int    `json:"seriesId"`
		SourceTitle string `json:"sourceTitle"`
		Quality     struct {
			Quality struct {
				Id   int    `json:"id"`
				Name string `json:"name"`
			} `json:"quality"`
			Revision struct {
				Version int `json:"version"`
				Real    int `json:"real"`
			} `json:"revision"`
		} `json:"quality"`
		QualityCutoffNotMet bool      `json:"qualityCutoffNotMet"`
		Date                time.Time `json:"date"`
		DownloadId          string    `json:"downloadId,omitempty"`
		EventType           string    `json:"eventType"`
		Data                struct {
			DroppedPath    string `json:"droppedPath,omitempty"`
			ImportedPath   string `json:"importedPath,omitempty"`
			DownloadClient string `json:"downloadClient,omitempty"`
			Reason         string `json:"reason,omitempty"`
		} `json:"data"`
		Episode struct {
			SeriesId                 int       `json:"seriesId"`
			EpisodeFileId            int       `json:"episodeFileId"`
			SeasonNumber             int       `json:"seasonNumber"`
			EpisodeNumber            int       `json:"episodeNumber"`
			Title                    string    `json:"title"`
			AirDate                  string    `json:"airDate"`
			AirDateUtc               time.Time `json:"airDateUtc"`
			Overview                 string    `json:"overview"`
			HasFile                  bool      `json:"hasFile"`
			Monitored                bool      `json:"monitored"`
			AbsoluteEpisodeNumber    int       `json:"absoluteEpisodeNumber"`
			UnverifiedSceneNumbering bool      `json:"unverifiedSceneNumbering"`
			Id                       int       `json:"id"`
		} `json:"episode"`
		Series struct {
			Title             string        `json:"title"`
			SortTitle         string        `json:"sortTitle"`
			SeasonCount       int           `json:"seasonCount"`
			Status            string        `json:"status"`
			Overview          string        `json:"overview"`
			Network           string        `json:"network"`
			AirTime           string        `json:"airTime"`
			Images            []interface{} `json:"images"`
			Seasons           []interface{} `json:"seasons"`
			Year              int           `json:"year"`
			Path              string        `json:"path"`
			ProfileId         int           `json:"profileId"`
			SeasonFolder      bool          `json:"seasonFolder"`
			Monitored         bool          `json:"monitored"`
			UseSceneNumbering bool          `json:"useSceneNumbering"`
			Runtime           int           `json:"runtime"`
			TvdbId            int           `json:"tvdbId"`
			TvRageId          int           `json:"tvRageId"`
			TvMazeId          int           `json:"tvMazeId"`
			FirstAired        time.Time     `json:"firstAired"`
			LastInfoSync      time.Time     `json:"lastInfoSync"`
			SeriesType        string        `json:"seriesType"`
			CleanTitle        string        `json:"cleanTitle"`
			ImdbId            string        `json:"imdbId"`
			TitleSlug         string        `json:"titleSlug"`
			Certification     string        `json:"certification"`
			Genres            []interface{} `json:"genres"`
			Tags              []int         `json:"tags"`
			Added             time.Time     `json:"added"`
			Ratings           struct {
				Votes int     `json:"votes"`
				Value float64 `json:"value"`
			} `json:"ratings"`
			QualityProfileId int `json:"qualityProfileId"`
			Id               int `json:"id"`
		} `json:"series"`
		Id int `json:"id"`
	} `json:"records"`
}
