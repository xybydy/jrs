package types

type MovieEditorEditBody struct {
	MovieIds            []int  `json:"movieIds"`
	Monitored           bool   `json:"monitored"`
	QualityProfileId    int    `json:"qualityProfileId"`
	MinimumAvailability string `json:"minimumAvailability"`
	RootFolderPath      string `json:"rootFolderPath"`
	Tags                []int  `json:"tags"`
	ApplyTags           string `json:"applyTags"`
	MoveFiles           bool   `json:"moveFiles"`
}

type MovieEditorRemoveBody struct {
	MovieIds           []int `json:"movieIds"`
	DeleteFIles        bool  `json:"deleteFIles"`
	AddImportExclusion bool  `json:"addImportExclusion"`
}
