package types

type Quality struct {
	Quality struct {
		Id         int    `json:"id"`
		Name       string `json:"name"`
		Source     string `json:"source"`
		Resolution int    `json:"resolution"`
		Modifier   string `json:"modifier"`
	} `json:"quality"`
	Revision Revision `json:"revision"`
}

type Revision struct {
	Version  int  `json:"version"`
	Real     int  `json:"real"`
	IsRepack bool `json:"isRepack"`
}

type QualityProfile struct {
	ID                int64     `json:"id"`
	Name              string    `json:"name"`
	UpgradeAllowed    bool      `json:"upgradeAllowed"`
	Cutoff            int64     `json:"cutoff"`
	Qualities         []Quality `json:"items"`
	MinFormatScore    int64     `json:"minFormatScore"`
	CutoffFormatScore int64     `json:"cutoffFormatScore"`
	FormatItems       []struct {
		Format int    `json:"format"`
		Name   string `json:"name"`
		Score  int    `json:"score"`
	} `json:"formatItems,omitempty"`
	Language Language `json:"language"`
}
