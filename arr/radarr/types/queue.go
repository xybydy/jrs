package types

import "time"

type Queue struct {
	Page          int            `json:"page"`
	PageSize      int            `json:"pageSize"`
	SortKey       string         `json:"sortKey"`
	SortDirection string         `json:"sortDirection"`
	TotalRecords  int            `json:"totalRecords"`
	Records       []*QueueRecord `json:"records"`
}

// QueueRecord is part of the activity Queue.
type QueueRecord struct {
	MovieID                 int64          `json:"movieId"`
	Languages               []Language     `json:"languages"`
	Quality                 Quality        `json:"quality"`
	CustomFormats           []CustomFormat `json:"customFormats"` // probably []int64
	Size                    float64        `json:"size"`
	Title                   string         `json:"title"`
	Sizeleft                float64        `json:"sizeleft"`
	Timeleft                string         `json:"timeleft"`
	EstimatedCompletionTime time.Time      `json:"estimatedCompletionTime"`
	Status                  string         `json:"status"`
	TrackedDownloadStatus   string         `json:"trackedDownloadStatus"`
	TrackedDownloadState    string         `json:"trackedDownloadState"`
	StatusMessages          []struct {
		Title    string   `json:"title"`
		Messages []string `json:"messages"`
	} `json:"statusMessages"`
	DownloadID     string `json:"downloadId"`
	Protocol       string `json:"protocol"`
	DownloadClient string `json:"downloadClient"`
	Indexer        string `json:"indexer"`
	OutputPath     string `json:"outputPath"`
	ID             int64  `json:"id"`
	ErrorMessage   string `json:"errorMessage"`
}
