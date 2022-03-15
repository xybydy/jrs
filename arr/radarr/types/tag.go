package types

type Tag struct {
	ID    int
	Label string
}

type TagDetail struct {
	Id              int    `json:"id"`
	Label           string `json:"label"`
	DelayProfileIds []int  `json:"delayProfileIds"`
	NotificationIds []int  `json:"notificationIds"`
	RestrictionIds  []int  `json:"restrictionIds"`
	NetImportIds    []int  `json:"netImportIds"`
	MovieIds        []int  `json:"movieIds"`
}
