package main

import (
	"flag"
	"log"
	"os"
	"text/template"
)

var configFilePath = flag.String("config", "config.json", "Path to config file.")
var doajXMLFilePath = flag.String("xml", "DOAJ.xml", "Path to DOAJ XML file.")
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

	t := template.Must(template.ParseFiles("crossrefOutput.tmpl"))
	err = t.Execute(os.Stdout, &templateData)
	if err != nil {
		log.Fatalln(err)
	}

}
