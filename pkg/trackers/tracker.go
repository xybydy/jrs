package trackers

import (
	"io"
	"net/http"
)

type Tracker interface {
	BuildRequest(method string, body io.Reader, args ...string) (*http.Request, error)
	GetIndexers() (*http.Request, error)
	GetIndexerSchema() (*http.Request, error)
	DeleteIndexer(i IndexerSchema) (*http.Request, error)
}
