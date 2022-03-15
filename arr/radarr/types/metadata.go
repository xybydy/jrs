package types

type Metadata struct {
	Enable bool   `json:"enable"`
	Name   string `json:"name"`
	Fields []struct {
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
