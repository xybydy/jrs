package types

type Collection struct {
	Name   string  `json:"name"`
	TmdbId int     `json:"tmdbId"`
	Images []Image `json:"images"`
}
