package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	loaded   bool
	Pwd      string
	Templays map[string]string `yaml:"templays"`
}

func (c *Config) SetPwd() error {
	if dir, err := filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		return fmt.Errorf("could not set pwd: %v", err)
	} else {
		c.Pwd = dir
	}
	return nil
}

func (c *Config) Load() error {
	templaysYAML, err := ioutil.ReadFile(filepath.Join(c.Pwd, ".templays.yml"))
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
		_, err := os.Stat(filepath.Join(c.Pwd, path))
		if os.IsNotExist(err) {
			return fmt.Errorf("templay %s does not exist", templay)
		}
	}
	return nil
}
