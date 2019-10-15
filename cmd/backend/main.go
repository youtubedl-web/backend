package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

var path string

func init() {
	// setup flags
	flag.StringVar(&path, "path", "config.json", "Path to config file")

	// setup logrus;

	// logs on json format
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})

	// include method info
	log.SetReportCaller(true)
}

func main() {
	c, err := loadConfig(path)
	if err != nil {
		log.Error(err)
	}

	// if the development mode is not active
	// change logrus level to warnings
	log.SetLevel(log.ErrorLevel)

	logPath := "backend.log"
	if len(c.Log) > 0 {
		logPath = c.Log
	}

	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Error(err)
	}

	log.SetOutput(f)
}
