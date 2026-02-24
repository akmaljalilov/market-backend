package utils

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func ReadYml(file string, value interface{}) error {
	yamlFile, err := ioutil.ReadFile(filepath.Clean(file))
	if err != nil {
		return err
	}
	r := bytes.NewReader(yamlFile)
	dec := yaml.NewDecoder(r)
	if err := dec.Decode(value); err != nil {
		return err
	}
	return nil
}
