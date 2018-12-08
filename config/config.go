package config

import (
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/burntsushi/toml"
)

const (
	DEFAULT_CONFIG_PATH = "config.toml"
)

var (
	ConfPath string
	Params   *Config
)

type Config struct {
	Dest []Destination `toml:"destinations"`
}

type Destination struct {
	Name string
	Path string
	Api  string
}

func init() {
	if Params == nil {
		Params = new(Config)
	}
}

func (c *Config) GetDestination(name string) *Destination {
	for _, i := range c.Dest {
		if strings.ToUpper(name) == strings.ToUpper(i.Name) {
			return &i
		}
	}
	return nil
}

func (c *Config) SaveFile(path string) error {
	ext := filepath.Ext(path)
	if ext == "" {
		path += ".toml"
	} else if ext != ".toml" {
		log.Fatal("Config file does not have toml extension, please correct it")
	}

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		return err
	}
	defer f.Close()

	if err := toml.NewEncoder(f).Encode(c); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func (c *Config) addDestination(name, param string, value interface{}) {
	dest := &Destination{name, "", ""}
	field := reflect.ValueOf(dest).Elem().FieldByName(strings.Title(param))

	if field.CanSet() {
		if field.Kind() == reflect.Uint16 {
			if reflect.TypeOf(value).String() == "int" {
				newValue := uint64(value.(int))
				field.SetUint(newValue)
			}
		}
		if field.Kind() == reflect.String {
			newValue := value.(string)
			field.SetString(newValue)
		}
	}
	c.Dest = append(c.Dest, *dest)
}

func (c *Config) ChangeParams(dest, param string, value interface{}) {
	main := reflect.ValueOf(c)
	mainElem := main.Elem()
	items := mainElem.FieldByName("Dest")

	if items.Len() > 0 {
		exists := false
		// Loops through already existed parameters to update them with the new parameters
		for i := 0; i < items.Len(); i++ {
			item := items.Index(i)
			if item.FieldByName("Name").String() == dest {
				exists = true
				field := item.FieldByName(param)
				if field.IsValid() {
					if field.CanSet() {
						if field.Kind() == reflect.Uint16 {
							if reflect.TypeOf(value).String() == "int" {
								newValue := uint64(value.(int))
								field.SetUint(newValue)
							}
						}
						if field.Kind() == reflect.String {
							newValue := value.(string)
							field.SetString(newValue)
						}
					}
				}
			}
		}
		if !exists {
			c.addDestination(dest, param, value)
		}
	} else {
		c.addDestination(dest, param, value)
	}

}

func ParseConfigFile() {
	if _, err := toml.DecodeFile(ConfPath, &Params); err != nil {
		log.Fatal(err)
	}
	if len(Params.Dest) < 1 {
		log.Fatal("There is no Destination configured.")
	}

	for _, dest := range Params.Dest {
		if dest.Api == "" {
			param := os.Getenv(strings.ToUpper(dest.Name))

			if param != "" {
				Params.GetDestination(dest.Name).Api = param
			} else {
				if len(Params.Dest) == 1 {
					log.Fatalf("There is no %s API configured.", dest.Name)
				} else {
					log.Printf("There is no %s API configured.", dest.Name)
				}
			}
		}

	}
}
