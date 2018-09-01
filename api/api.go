package api

import "net/http"

type Indexers interface {
	ConfigIndexer() (*http.Request, error)
	GetIndexers() (*http.Request, error)
	GetIndexerSchema() (*http.Request, error)
	SetIndexer(p Param) (*http.Request, error)
}

type Param struct {
	Key   string
	Value string
}
