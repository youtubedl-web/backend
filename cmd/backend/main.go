package main

import (
	"flag"

	log "github.com/sirupsen/logrus"
)

var path string

func init() {
	// setup flags
	flag.StringVar(&path, "path", "config.json", "Path to config file")

	// setup logrus;
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})
}

func main() {
	_, err := loadConfig(path)
	if err != nil {
		log.Error(err)
	}
}
