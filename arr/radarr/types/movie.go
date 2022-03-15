package types

type Movie struct {
	Id                  int        `json:"id"`
	Title               string     `json:"title"`
	SortTitle           string     `json:"sortTitle"`
	SizeOnDisk          int        `json:"sizeOnDisk"`
	Overview            string     `json:"overview"`
	InCinemas           string     `json:"inCinemas"`
	PhysicalRelease     string     `json:"physicalRelease"`
	Images              []Image    `json:"images"`
	Website             string     `json:"website"`
	Year                int        `json:"year"`
	HasFile             bool       `json:"hasFile"`
	YouTubeTrailerId    string     `json:"youTubeTrailerId"`
	Studio              string     `json:"studio"`
	Path                string     `json:"path"`
	RootFolderPath      string     `json:"rootFolderPath"`
	QualityProfileId    int        `json:"qualityProfileId"`
	Monitored           bool       `json:"monitored"`
	MinimumAvailability string     `json:"minimumAvailability"`
	IsAvailable         bool       `json:"isAvailable"`
	FolderName          string     `json:"folderName"`
	Runtime             int        `json:"runtime"`
	CleanTitle          string     `json:"cleanTitle"`
	ImdbId              string     `json:"imdbId"`
	TmdbId              int        `json:"tmdbId"`
	TitleSlug           string     `json:"titleSlug"`
	Certification       string     `json:"certification"`
	Genres              []string   `json:"genres"`
	Tags                []int      `json:"tags"`
	Added               string     `json:"added"`
	Ratings             Rating     `json:"ratings"`
	Collection          Collection `json:"collection"`
	Status              string     `json:"status"`
}

type BulkMovieEdit struct {
	MovieIds            []int  `json:"movieIds"`
	Monitored           bool   `json:"monitored"`
	QualityProfileId    int    `json:"qualityProfileId"`
	MinimumAvailability string `json:"minimumAvailability"`
	RootFolderPath      string `json:"rootFolderPath"`
	Tags                []int  `json:"tags"`
	ApplyTags           string `json:"applyTags"`
	MoveFiles           bool   `json:"moveFiles"`
}

type BulkMovieDelete struct {
	MovieIds           []int `json:"movieIds"`
	DeleteFIles        bool  `json:"deleteFIles"`
	AddImportExclusion bool  `json:"addImportExclusion"`
}

type MediaInfo struct {
	AudioAdditionalFeatures string `json:"audioAdditionalFeatures"`
	AudioBitrate            int    `json:"audioBitrate"`
	AudioChannels           int    `json:"audioChannels"`
	AudioCodec              string `json:"audioCodec"`
	AudioLanguages          string `json:"audioLanguages"`
	AudioStreamCount        int    `json:"audioStreamCount"`
	VideoBitDepth           int    `json:"videoBitDepth"`
	VideoBitrate            int    `json:"videoBitrate"`
	VideoCodec              string `json:"videoCodec"`
	VideoFps                int    `json:"videoFps"`
	Resolution              string `json:"resolution"`
	RunTime                 string `json:"runTime"`
	ScanType                string `json:"scanType"`
	Subtitles               string `json:"subtitles"`
}

type MovieFile struct {
	MovieId             int        `json:"movieId"`
	RelativePath        string     `json:"relativePath"`
	Path                string     `json:"path"`
	Size                int        `json:"size"`
	DateAdded           string     `json:"dateAdded"`
	IndexerFlags        int        `json:"indexerFlags"`
	Quality             Quality    `json:"quality"`
	MediaInfo           MediaInfo  `json:"mediaInfo"`
	QualityCutoffNotMet bool       `json:"qualityCutoffNotMet"`
	Languages           []Language `json:"languages"`
	ReleaseGroup        string     `json:"releaseGroup"`
	Id                  int        `json:"id"`
}
