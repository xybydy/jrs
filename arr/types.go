package arr

type Image struct {
	CoverType string
	Url       string
}

type IndexerSchema struct {
	EnableRss     bool   `json:"enableRss"`
	EnableSearch  bool   `json:"enableSearch"`
	SupportRss    bool   `json:"supportsRss"`
	SupportSearch bool   `json:"supportsSearch"`
	Protocol      string `json:"protocol"`
	Name          string `json:"name"`
	Fields        []struct {
		Order         int         `json:"order"`
		Name          string      `json:"name"`
		Label         string      `json:"label"`
		Tip           string      `json:"type"`
		Advanced      bool        `json:"advanced"`
		HelpText      string      `json:"helpText,omitempty"`
		HelpLink      string      `json:"helpLink,omitempty"`
		Value         interface{} `json:"value,omitempty"`
		SelectOptions []struct {
			Value int    `json:"value"`
			Name  string `json:"name"`
		} `json:"selectOptions,omitempty"`
	} `json:"fields,omitempty"`
	ID                 int             `json:"id,omitempty"`
	ImplementationName string          `json:"implementationName"`
	Implementation     string          `json:"implementation"`
	ConfigContract     string          `json:"configContract"`
	InfoLink           string          `json:"infoLink"`
	Presets            []IndexerSchema `json:"presets"`
}

type IndexerSchemas []IndexerSchema

func (i *IndexerSchemas) GetTorznab() *IndexerSchema {
	for _, k := range *i {
		if k.ImplementationName == "Torznab" && k.Protocol == "torrent" && k.Name == "" {
			return &k
		}
	}
	return nil
}
