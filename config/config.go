package config

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/gcfg.v1"
)

// ReadConfig is file handler for reading configuration files into variable
// Param: -  filePath string
// Return: - boolean
func (cfg *Config) ReadConfig(filePath string) bool {

	var configString []string

	config, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("[Err] in reading Config File: %+v", err)
		return false
	}

	configString = append(configString, string(config))

	err = gcfg.ReadStringInto(cfg, strings.Join(configString, "\n\n"))
	if err != nil {
		log.Printf("[Err] in reading Config: %+v", err)
		return false
	}

	return true
}
