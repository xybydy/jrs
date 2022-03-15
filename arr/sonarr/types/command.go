package types

import "time"

type Command struct {
	Name                string    `json:"name"`
	StartedOn           time.Time `json:"startedOn"`
	StateChangeTime     time.Time `json:"stateChangeTime"`
	SendUpdatesToClient bool      `json:"sendUpdatesToClient"`
	State               string    `json:"state"`
	Id                  int       `json:"id"`
}
