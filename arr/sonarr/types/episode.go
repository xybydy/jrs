package types

import "time"

type Episode struct {
	SeriesId              int       `json:"seriesId"`
	EpisodeFileId         int       `json:"episodeFileId"`
	SeasonNumber          int       `json:"seasonNumber"`
	EpisodeNumber         int       `json:"episodeNumber"`
	Title                 string    `json:"title"`
	AirDate               string    `json:"airDate"`
	AirDateUtc            time.Time `json:"airDateUtc"`
	Overview              string    `json:"overview"`
	HasFile               bool      `json:"hasFile"`
	Monitored             bool      `json:"monitored"`
	SceneEpisodeNumber    int       `json:"sceneEpisodeNumber"`
	SceneSeasonNumber     int       `json:"sceneSeasonNumber"`
	TvDbEpisodeId         int       `json:"tvDbEpisodeId"`
	AbsoluteEpisodeNumber int       `json:"absoluteEpisodeNumber"`
	Id                    int       `json:"id"`
}
