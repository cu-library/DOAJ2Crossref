package main

import (
	"flag"
	"log"
)

var configFilePath = flag.String("config", "config.json", "Path to config file.")

func main() {
	flag.Parse()

	mappingConfig, err := LoadMappingConfig(*configFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("%v", mappingConfig)
}
