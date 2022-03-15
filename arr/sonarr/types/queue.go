package types

import "time"

type Queue struct {
	Series  Series  `json:"series"`
	Episode Episode `json:"episode"`
	Quality struct {
		Quality struct {
			Id   int    `json:"id"`
			Name string `json:"name"`
		} `json:"quality"`
		Revision struct {
			Version int `json:"version"`
			Real    int `json:"real"`
		} `json:"revision"`
	} `json:"quality"`
	Size                    int64         `json:"size"`
	Title                   string        `json:"title"`
	Sizeleft                int           `json:"sizeleft"`
	Timeleft                string        `json:"timeleft"`
	EstimatedCompletionTime time.Time     `json:"estimatedCompletionTime"`
	Status                  string        `json:"status"`
	TrackedDownloadStatus   string        `json:"trackedDownloadStatus"`
	StatusMessages          []interface{} `json:"statusMessages"`
	DownloadId              string        `json:"downloadId"`
	Protocol                string        `json:"protocol"`
	Id                      int           `json:"id"`
}
