package logger

import (
	"log"
	"os"
	"path/filepath"
)

//InitLogger : creates log File in append mode and inits logger
func InitLogger(logfilePath string) error {

	//create filepath
	err := os.MkdirAll(filepath.Dir(logfilePath), 0755)
	if err != nil && err != os.ErrExist {
		log.Printf("[Err] error in creating file:%s err: %v", logfilePath, err)
		return err
	}

	//open log file in append mode
	logFile, err := os.OpenFile(logfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("[Err] error in opening file: %v", err)
		return err
	}

	//set logFile as destination for log
	log.SetOutput(logFile)

	return nil
}
