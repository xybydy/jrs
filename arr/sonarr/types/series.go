package types

import "time"

type Series struct {
	TvdbId           int       `json:"tvdbId"`
	TvRageId         int       `json:"tvRageId"`
	ImdbId           string    `json:"imdbId"`
	Title            string    `json:"title"`
	CleanTitle       string    `json:"cleanTitle"`
	Status           string    `json:"status"`
	Overview         string    `json:"overview"`
	AirTime          string    `json:"airTime"`
	Monitored        bool      `json:"monitored"`
	QualityProfileId int       `json:"qualityProfileId"`
	SeasonFolder     bool      `json:"seasonFolder"`
	LastInfoSync     time.Time `json:"lastInfoSync"`
	Runtime          int       `json:"runtime"`
	Images           []struct {
		CoverType string `json:"coverType"`
		Url       string `json:"url"`
	} `json:"images"`
	SeriesType        string    `json:"seriesType"`
	Network           string    `json:"network"`
	UseSceneNumbering bool      `json:"useSceneNumbering"`
	TitleSlug         string    `json:"titleSlug"`
	Path              string    `json:"path"`
	Year              int       `json:"year"`
	FirstAired        time.Time `json:"firstAired"`
	QualityProfile    struct {
		Value struct {
			Name    string `json:"name"`
			Allowed []struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Weight int    `json:"weight"`
			} `json:"allowed"`
			Cutoff struct {
				Id     int    `json:"id"`
				Name   string `json:"name"`
				Weight int    `json:"weight"`
			} `json:"cutoff"`
			Id int `json:"id"`
		} `json:"value"`
		IsLoaded bool `json:"isLoaded"`
	} `json:"qualityProfile"`
	Seasons []struct {
		SeasonNumber int  `json:"seasonNumber"`
		Monitored    bool `json:"monitored"`
	} `json:"seasons"`
	Id int `json:"id"`
}
