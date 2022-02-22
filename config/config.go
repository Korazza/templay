package config

import (
	"fmt"
	"os"

	"github.com/Korazza/templay/utils"
)

const (
	TEMPLAY_CONFIGFILE = ".templays.yaml"
	TEMPLAY_EXTENSION  = ".tp"
)

type Config struct {
	loaded   bool
	Templays map[string]string `yaml:"templays"`
}

func (c *Config) Load() error {
	f, err := os.Open(TEMPLAY_CONFIGFILE)
	if err != nil {
		return fmt.Errorf("failed to open configuration file %v", TEMPLAY_CONFIGFILE)
	}
	defer f.Close()

	err = utils.ParseYaml(f, c)
	if err != nil {
		return fmt.Errorf("failed to parse configuration file %v", TEMPLAY_CONFIGFILE)
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
