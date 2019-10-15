package main

import (
	"encoding/json"
	"os"
)

type config struct {
	Development bool

	// Paths
	YoutubeExecutable string
	Log               string

	// API web server
	Host string
	Port int
}

func loadConfig(path string) (*config, error) {
	cfg := &config{}

	cfgPath := "config.json"
	if len(path) > 0 {
		cfgPath = path
	}

	f, err := os.Open(cfgPath)
	if err != nil {
		return nil, err
	}

	parser := json.NewDecoder(f)
	err = parser.Decode(&cfg)

	return cfg, nil
}
