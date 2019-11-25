package main

import (
	"flag"
	"log"
	"os"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"

	"github.com/youtubedl-web/backend"
	h "github.com/youtubedl-web/backend/http"
)

var path string
var logger = logrus.New()

func init() {
	// setup flags
	flag.StringVar(&path, "path", "config.json", "Path to config file")
	flag.Parse()

	// setup logrus

	// logs on json format
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})

	// include method info
	logger.SetReportCaller(true)
}

func main() {
	c, err := loadConfig(path)
	if err != nil {
		color.Red("Couldn't open JSON config file")
		os.Exit(1)
	}

	cfg := &backend.Config{
		Development:    c.Development,
		Host:           c.Host,
		Port:           c.Port,
		Logger:         logger,
		ExecutablePath: c.YoutubeExecutable,
		Storage:        c.Storage,
		PublicHost:     c.PublicHost,
	}

	// if the development mode is not active
	// change logrus level to warnings
	logger.SetLevel(logrus.ErrorLevel)

	// open log output file
	logPath := "backend.log"
	if len(c.Log) > 0 {
		logPath = c.Log
	}

	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		color.Red("Couldn't open log file (filename: %s)", logPath)
		os.Exit(1)
	}

	// set logs output file
	log.SetOutput(f)

	h.Serve(cfg)
}
