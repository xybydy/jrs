package types

type CustomFormat struct {
	Id                              int             `json:"id,omitempty"`
	Name                            string          `json:"name"`
	IncludeCustomFormatWhenRenaming bool            `json:"includeCustomFormatWhenRenaming"`
	Specifications                  []Specification `json:"specifications"`
}
