package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// Config holds data from the json config file.
type Config struct {
	Mappings []struct {
		JournalTitle            string `json:"journalTitle"`
		Prefix                  string `json:"prefix"`
		AbbreviatedJournalTitle string `json:"abbreviatedJournalTitle"`
	} `json:"mappings"`
	Orcids []struct {
		Name  string `json:"name"`
		Orcid string `json:"orcid"`
	} `json:"orcids"`
}

// PrefixAndAbbreviation is a tuple which holds a prefix and an abbreviation for a journal title.
type PrefixAndAbbreviation struct {
	Prefix       string
	Abbreviation string
}

// LoadConfig returns a pointer to a Config loaded with mappings from the config file.
func LoadConfig(configFilePath string) (map[string]PrefixAndAbbreviation, map[string]string, error) {

	config := new(Config)
	mapping := make(map[string]PrefixAndAbbreviation)
	orcids := make(map[string]string)

	absoluteConfigFilePath, err := filepath.Abs(configFilePath)

	configFile, err := os.Open(absoluteConfigFilePath)
	if err != nil {
		return mapping, orcids, err
	}
	defer configFile.Close()

	configDecoder := json.NewDecoder(configFile)
	err = configDecoder.Decode(config)
	if err != nil {
		return mapping, orcids, err
	}

	for _, configMapping := range config.Mappings {
		mapping[configMapping.JournalTitle] = PrefixAndAbbreviation{configMapping.Prefix, configMapping.AbbreviatedJournalTitle}
	}

	for _, orcidpair := range config.Orcids {
		orcids[orcidpair.Name] = orcidpair.Orcid
	}

	return mapping, orcids, nil
}
