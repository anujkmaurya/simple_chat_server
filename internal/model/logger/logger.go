package logger

import (
	"log"
	"os"
	"path/filepath"
)

func InitLogger(logfilePath string) error {

	err := os.MkdirAll(filepath.Dir(logfilePath), 0755)
	if err != nil && err != os.ErrExist {
		log.Printf("error in creating file:%s err: %v", logfilePath, err)
		return err
	}

	logFile, err := os.OpenFile(logfilePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("error in opening file: %v", err)
		return err
	}
	// defer logFile.Close()
	log.SetOutput(logFile)
	return nil
}
