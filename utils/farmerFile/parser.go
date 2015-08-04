package farmerFile

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type ConfigFile struct {
	Image   string
	Ports   []string
	Env     []string
	Scripts map[string]string
}

func Parse(address string) (ConfigFile, error) {
	filename, _ := filepath.Abs(address + "/.farmer.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		return ConfigFile{}, err
	}

	var farmerFile ConfigFile
	err = yaml.Unmarshal(yamlFile, &farmerFile)

	if err != nil {
		return ConfigFile{}, err
	}

	return farmerFile, nil
}
