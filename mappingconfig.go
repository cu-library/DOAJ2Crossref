package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type mappingFromConfig struct {
	Mappings []struct {
		JournalTitle string `json:"journalTitle"`
		Prefix       string `json:"prefix"`
	} `json:"mappings"`
}

// LoadMappingConfig returns a pointer to a MappingConfig loaded with mappings from the config file.
func LoadMappingConfig(configFilePath string) (map[string]string, error) {

	config := new(mappingFromConfig)
	mapping := make(map[string]string)

	absoluteConfigFilePath, err := filepath.Abs(configFilePath)

	configFile, err := os.Open(absoluteConfigFilePath)
	if err != nil {
		return mapping, err
	}
	defer configFile.Close()

	configDecoder := json.NewDecoder(configFile)
	err = configDecoder.Decode(config)
	if err != nil {
		return mapping, err
	}

	for _, configMapping := range config.Mappings {
		mapping[configMapping.JournalTitle] = configMapping.Prefix
	}

	return mapping, nil
}
