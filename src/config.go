package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type ConfigItem struct {
	Repository string
	Password   string
}

type Config map[string]ConfigItem

func readConfig(file string) map[string]ConfigItem {
	config := map[string]ConfigItem{}
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return config
}
