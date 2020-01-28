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

	// Public Address
	AllowedOrigins            string
	AllowedOriginsDevelopment string
	PublicHost                string

	// API web server
	Secure bool
	Host   string
	Port   int

	Storage string
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
