package main

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

func LoadDOAJ(xmlFilePath string) (*DOAJRecords, error) {

	records := new(DOAJRecords)
	absoluteXMLFilePath, err := filepath.Abs(xmlFilePath)

	xmlFile, err := os.Open(absoluteXMLFilePath)
	if err != nil {
		return records, err
	}
	defer xmlFile.Close()

	xmlDecoder := xml.NewDecoder(xmlFile)
	err = xmlDecoder.Decode(records)
	if err != nil {
		return records, err
	}

	return records, nil
}
