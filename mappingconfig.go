package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type MappingFromConfig struct {
	Mappings []struct {
		JournalTitle string `json:"journalTitle"`
		Prefix       string `json:"prefix"`
		AbbreviatedJournalTitle string `json:"abbreviatedJournalTitle"`
	} `json:"mappings"`
}

type PrefixAndAbbreviation struct {
	Prefix string
	Abbreviation string
}

// LoadMappingConfig returns a pointer to a MappingConfig loaded with mappings from the config file.
func LoadMappingConfig(configFilePath string) (map[string]PrefixAndAbbreviation, error) {

	config := new(MappingFromConfig)
	mapping := make(map[string]PrefixAndAbbreviation)

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
		mapping[configMapping.JournalTitle] = PrefixAndAbbreviation{configMapping.Prefix, configMapping.AbbreviatedJournalTitle}
	}

	return mapping, nil
}
