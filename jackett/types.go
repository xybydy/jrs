package jackett

type Caps struct {
	ID   string `json:"ID"`
	Name string `json:"Name"`
}

type Indexer struct {
	ID               string
	Name             string
	Description      string
	Type             string
	Configured       bool
	SiteLink         string `json:"site_link"`
	Alternativelinks []string
	Language         string
	LastError        string `json:"last_error"`
	Potatoenabled    bool
	Caps             []Caps
}

type IndexerConfig []struct {
	ID      string
	Type    string            `json:",omitempty"`
	Name    string            `json:",omitempty"`
	Value   string            `json:",omitempty"`
	Options map[string]string `json:",omitempty"`
}

func (ic *IndexerConfig) UpdateField(id, param string) {
	for _, i := range *ic {
		if i.ID == id {
			i.Value = param
		}
	}
}

func (ic *IndexerConfig) SetCredentials(user, passwd string) {
	ic.UpdateField("username", user)
	ic.UpdateField("password", passwd)
}

type Indexers []Indexer

func (i *Indexers) GetIndexer(id string) *Indexer {
	for _, k := range *i {
		if k.ID == id {
			return &k
		}
	}
	return nil
}
