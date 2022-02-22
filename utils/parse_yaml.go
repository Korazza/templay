package utils

import (
	"io"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func ParseYaml(reader io.Reader, out interface{}) error {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(b, out)
	if err != nil {
		return err
	}

	return nil
}
