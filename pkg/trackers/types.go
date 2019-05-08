package trackers

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
type CommandSummaryResult struct {
	Name    string `json:"name,omitempty"`
	Message string `json:"message,omitempty"`
	Body    struct {
		MovieIds            []int  `json:"movieIds,omitempty"`
		SendUpdatesToClient bool   `json:"sendUpdatesToClient,omitempty"`
		UpdateScheduledTask bool   `json:"updateScheduledTask,omitempty"`
		CompletionMessage   string `json:"completionMessage,omitempty"`
		Name                string `json:"name,omitempty"`
		Trigger             string `json:"trigger,omitempty"`
	} `json:"body,omitempty"`
	Priority            string `json:"priority,omitempty"`
	Status              string `json:"status,omitempty"`
	Queued              string `json:"queued,omitempty"`  // Date
	Started             string `json:"started,omitempty"` // Date
	Trigger             string `json:"trigger,omitempty"`
	State               string `json:"state,omitempty"`
	Manual              bool   `json:"manual,omitempty"`
	StartedOn           string `json:"startedOn,omitempty"`       // Date
	StateChangeTime     string `json:"stateChangeTime,omitempty"` //Date
	SendUpdatesToClient bool   `json:"sendUpdatesToClient,omitempty"`
	UpdateScheduledTask bool   `json:"updateScheduledTask,omitempty"`
	Id                  int    `json:"id,omitempty"`
}
