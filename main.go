package main

import (
	"flag"
	"log"
)

var configFilePath = flag.String("config", "config.json", "Path to config file.")
var doajXMLFilePath = flag.String("xml", "DOAJ.xml", "Path to DOAJ XML file.")

func main() {
	flag.Parse()

	mappingConfig, err := LoadMappingConfig(*configFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(mappingConfig)

	doajData, err := LoadDOAJ(*doajXMLFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, record := range doajData.DOAJRecords {
		log.Println("%v", record.DOAJAbstract)
		log.Println("%v", record.DOAJAffiliationsList)
		log.Println("%v", record.DOAJAuthors)
		log.Println("%v", record.DOAJDocumentType)
		log.Println("%v", record.DOAJDoi)
		log.Println("%v", record.DOAJEndPage)
		log.Println("%v", record.DOAJFullTextUrl)
		log.Println("%v", record.DOAJIssn)
		log.Println("%v", record.DOAJIssue)
		log.Println("%v", record.DOAJJournalTitle)
		log.Println("%v", record.DOAJKeywords)
		log.Println("%v", record.DOAJLanguage)
		log.Println("%v", record.DOAJPublicationDate)
		log.Println("%v", record.DOAJPublisher)
		log.Println("%v", record.DOAJPublisherRecordId)
		log.Println("%v", record.DOAJStartPage)
		log.Println("%v", record.DOAJTitle)
		log.Println("%v", record.DOAJVolume)
	}

}
