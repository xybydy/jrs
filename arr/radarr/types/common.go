package types

type Language struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Specification struct {
	Name               string `json:"name"`
	Implementation     string `json:"implementation"`
	ImplementationName string `json:"implementationName"`
	InfoLink           string `json:"infoLink"`
	Negate             bool   `json:"negate"`
	Required           bool   `json:"required"`
	Fields             []struct {
		Order    int    `json:"order"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		HelpText string `json:"helpText"`
		Type     string `json:"type"`
		Advanced bool   `json:"advanced"`
	}
}
