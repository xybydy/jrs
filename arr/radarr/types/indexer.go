package types

type Indexer struct {
	EnableRss               bool   `json:"enableRss"`
	EnableAutomaticSearch   bool   `json:"enableAutomaticSearch"`
	EnableInteractiveSearch bool   `json:"enableInteractiveSearch"`
	SupportsRss             bool   `json:"supportsRss"`
	SupportsSearch          bool   `json:"supportsSearch"`
	Protocol                string `json:"protocol"`
	Priority                int    `json:"priority"`
	Name                    string `json:"name"`
	Fields                  []struct {
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
	Tags               []Tag  `json:"tags"`
	Id                 int    `json:"id"`
}
