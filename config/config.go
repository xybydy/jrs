package config

import "strings"

type Config struct {
	Jackett Jackett
	Dest    []Destination `toml:"destinations"`
}

type Destination struct {
	Name string
	Ip   string
	Port int
	Api  string
}

type Jackett struct {
	Ip   string
	Port int
	Api  string
}

func (c *Config) GetDestination(name string) *Destination {
	for _, i := range c.Dest {
		if strings.ToUpper(name) == strings.ToUpper(i.Name) {
			return &i
		}
	}
	return nil
}
