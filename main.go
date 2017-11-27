package main

import (
	"encoding/csv"
	"flag"
	"log"
	"os"
	"text/template"
)

var configFilePath = flag.String("config", "config.json", "Path to config file.")
var doajXMLFilePath = flag.String("in", "DOAJ.xml", "Path to DOAJ XML file.")
var crossrefOutputFilePath = flag.String("out", "crossref.xml", "Path to which the output XML file will be written.")
var urlToDOICSVOutputFilePath = flag.String("report", "report.csv", "Path to which the report csv file will be written.")
var depositorName = flag.String("depositor", "", "Name of the organization registering the DOIs. The name placed in this element should match the name under which a depositing organization has registered with CrossRef.")
var depositorEmail = flag.String("email", "", "Email address to which batch success and/or error messages are sent. It is recommended that this address be unique to a position within the organization submitting data (e.g. \"doi@...\") rather than unique to a person. In this way, the alias for delivery of this mail can be changed as responsibility for submission of DOI data within the organization changes from one person to another.")
var registrant = flag.String("registrant", "", "The organization that owns the information being registered.")

func main() {
	flag.Parse()

	if *depositorName == "" {
		log.Fatalln("depositor required")
	}
	if *depositorEmail == "" {
		log.Fatalln("email required")
	}
	if *registrant == "" {
		log.Fatalln("registrant required")
	}

	mappingConfig, err := LoadMappingConfig(*configFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	doajData, err := LoadDOAJ(*doajXMLFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	templateData := CreateTemplateData(*depositorName, *depositorEmail, *registrant, mappingConfig, doajData)

	output, err := os.Create(*crossrefOutputFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer output.Close()

	report, err := os.Create(*urlToDOICSVOutputFilePath)
	if err != nil {
		log.Fatalln(err)
	}
	defer report.Close()

	t := template.Must(template.New("template").Parse(templateSkeleton))
	err = t.Execute(output, &templateData)
	if err != nil {
		log.Fatalln(err)
	}

	w := csv.NewWriter(report)

	err = w.Write([]string{"URI", "DOI"})
	if err != nil {
		log.Fatalln("Error writing to csv:", err)
	}

	for _, journal := range templateData.Journals {
		for _, article := range journal.Articles {
			err = w.Write([]string{article.URI, article.DOI})
			if err != nil {
				log.Fatalln("Error writing to csv:", err)
			}
		}

	}

	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatalln(err)
	}

}
