package sonarr

type IndexerSchema struct {
	EnableRss     bool `json:"enableRss"`
	EnableSearch  bool `json:"enableSearch"`
	SupportRss    bool `json:"supportsRss"`
	SupportSearch bool `json:"supportsSearch"`
	Protocol      string
	Name          string
	Fields        []struct {
		Order    int
		Name     string
		Label    string
		HelpText string      `json:"helpText,omitempty"`
		Value    interface{} `json:",omitempty"`
		Tip      string      `json:"type"`
		Advanced bool        `json:",omitempty"`
	}
	ID                 int    `json:"id,omitempty"`
	ImplementationName string `json:"implementationName"`
	Implementation     string
	ConfigContract     string `json:"configContract"`
	InfoLink           string `json:"infoLink"`
	Presets            []IndexerSchema
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
