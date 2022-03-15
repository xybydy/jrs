package types

type Notification struct {
	OnGrab                bool   `json:"onGrab"`
	OnDownload            bool   `json:"onDownload"`
	OnUpgrade             bool   `json:"onUpgrade"`
	OnRename              bool   `json:"onRename"`
	OnDelete              bool   `json:"onDelete"`
	OnHealthIssue         bool   `json:"onHealthIssue"`
	SupportsOnGrab        bool   `json:"supportsOnGrab"`
	SupportsOnDownload    bool   `json:"supportsOnDownload"`
	SupportsOnUpgrade     bool   `json:"supportsOnUpgrade"`
	SupportsOnRename      bool   `json:"supportsOnRename"`
	SupportsOnDelete      bool   `json:"supportsOnDelete"`
	SupportsOnHealthIssue bool   `json:"supportsOnHealthIssue"`
	IncludeHealthWarnings bool   `json:"includeHealthWarnings"`
	Name                  string `json:"name"`
	Fields                []struct {
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
	Message            struct {
		Message string `json:"message"`
		Type    string `json:"type"`
	} `json:"message"`
	Tags []int `json:"tags"`
	Id   int   `json:"id"`
}
