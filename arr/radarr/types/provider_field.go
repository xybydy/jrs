package types

type ProviderField struct {
	Order    int    `json:"order"`
	Name     string `json:"name"`
	Label    string `json:"label"`
	HelpText string `json:"helpText"`
	Type     string `json:"type"`
	Advanced bool   `json:"advanced"`
}
