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

// func (i *IndexerSchema) UpdateFields(p ...api.Param) {
// 	for _, k := range p {
// 		index := strings.Index(k.Key, ";")
// 		if index != -1 {
// 			r := reflect.ValueOf(i)
// 			f := reflect.Indirect(r).FieldByName(k.Key[:index])
// 			fmt.Printf()
// 		}
// 	}
// }

type IndexerSchemas []IndexerSchema

func (i *IndexerSchemas) GetTorznab() *IndexerSchema {
	for _, k := range *i {
		if k.ImplementationName == "Torznab" && k.Protocol == "torrent" && k.Name == "" {
			return &k
		}
	}
	return nil
}
