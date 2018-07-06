package config

import "strings"

type Config struct {
	Jackett Jackett
	Dest    []Destination `toml:"destinations"`
}

type Destination struct {
	Name string
	IP   string
	Port int
	API  string
}

type Jackett struct {
	IP   string
	Port int
	API  string
}

func (c *Config) GetDestination(name string) *Destination {
	for _, i := range c.Dest {
		if strings.ToUpper(name) == strings.ToUpper(i.Name) {
			return &i
		}
	}
	return nil
}
