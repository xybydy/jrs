package types

type BlockListRecord struct {
	MovieId       int            `json:"movieId"`
	SourceTitle   string         `json:"sourceTitle"`
	Languages     []Language     `json:"languages"`
	Quality       Quality        `json:"quality"`
	CustomFormats []CustomFormat `json:"customFormats"`
	Date          string         `json:"date"`
	Protocol      string         `json:"protocol"`
	Indexer       string         `json:"indexer"`
	Message       string         `json:"message"`
	Id            int            `json:"id"`
}
type Blocklist struct {
	Page          int               `json:"page"`
	PageSize      int               `json:"pageSize"`
	SortDirection string            `json:"sortDirection"`
	SortKey       string            `json:"sortKey"`
	TotalRecords  int               `json:"totalRecords"`
	Records       []BlockListRecord `json:"records"`
}
