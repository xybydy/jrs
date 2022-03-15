package types

type ImportList struct {
	Enabled             bool   `json:"enabled"`
	EnableAuto          bool   `json:"enableAuto"`
	ShouldMonitor       bool   `json:"shouldMonitor"`
	RootFolderPath      string `json:"rootFolderPath"`
	QualityProfileId    int    `json:"qualityProfileId"`
	SearchOnAdd         bool   `json:"searchOnAdd"`
	MinimumAvailability string `json:"minimumAvailability"`
	ListType            string `json:"listType"`
	ListOrder           int    `json:"listOrder"`
	Name                string `json:"name"`
	Fields              []struct {
		Order    int    `json:"order"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		HelpText string `json:"helpText"`
		Value    string `json:"value"`
		Type     string `json:"type"`
		Advanced bool   `json:"advanced"`
	} `json:"fields"`
	ImplementationName string `json:"implementationName"`
	Implementation     string `json:"implementation"`
	ConfigContract     string `json:"configContract"`
	InfoLink           string `json:"infoLink"`
	Tags               []int  `json:"tags"`
	Id                 int    `json:"id"`
}
