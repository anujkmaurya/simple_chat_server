package config

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/gcfg.v1"
)

// ReadConfig is file handler for reading configuration files into variable
// Param: -  filepath string
// Return: - boolean
func (cfg *Config) ReadConfig(fileName string) bool {

	var configString []string

	config, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Printf(" function ReadConfig %+v", err)
		return false
	}

	configString = append(configString, string(config))

	err = gcfg.ReadStringInto(cfg, strings.Join(configString, "\n\n"))
	if err != nil {
		log.Printf("function ReadConfig", err)
		return false
	}

	return true
}
