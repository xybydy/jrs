package types

import "time"

type EpisodeFile struct {
	SeriesId     int       `json:"seriesId"`
	SeasonNumber int       `json:"seasonNumber"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	DateAdded    time.Time `json:"dateAdded"`
	SceneName    string    `json:"sceneName"`
	Quality      struct {
		Quality struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"quality"`
		Proper bool `json:"proper"`
	} `json:"quality"`
	Id int `json:"id"`
}
