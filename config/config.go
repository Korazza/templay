package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const TEMPLAY_EXTENSION = ".tp"

type Config struct {
	loaded   bool
	Templays map[string]string `yaml:"templays"`
}

func (c *Config) Load() error {
	templaysYAML, err := ioutil.ReadFile(".templays.yml")
	if err != nil {
		return fmt.Errorf("no configuration file detected")
	}
	err = yaml.Unmarshal(templaysYAML, c)
	if err != nil {
		return err
	}
	c.loaded = true
	return nil
}

func (c *Config) Validate() error {
	if !c.loaded {
		return nil
	}
	for templay, path := range c.Templays {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			return fmt.Errorf("templay %s does not exist", templay)
		}
	}
	return nil
}
