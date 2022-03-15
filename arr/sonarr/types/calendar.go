package types

import "time"

type Calendar struct {
	SeriesId           int       `json:"seriesId"`
	EpisodeFileId      int       `json:"episodeFileId"`
	SeasonNumber       int       `json:"seasonNumber"`
	EpisodeNumber      int       `json:"episodeNumber"`
	Title              string    `json:"title"`
	AirDate            string    `json:"airDate"`
	AirDateUtc         time.Time `json:"airDateUtc"`
	Overview           string    `json:"overview"`
	HasFile            bool      `json:"hasFile"`
	Monitored          bool      `json:"monitored"`
	SceneEpisodeNumber int       `json:"sceneEpisodeNumber"`
	SceneSeasonNumber  int       `json:"sceneSeasonNumber"`
	TvDbEpisodeId      int       `json:"tvDbEpisodeId"`
	Series             Series    `json:"series"`
	Downloading        bool      `json:"downloading"`
	Id                 int       `json:"id"`
}
