package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileConfig struct {
	LocalFile     string `json:"local_file"`
	DriveFolderID string `json:"drive_folder_id"`
	DriveFileName string `json:"drive_file_name"`
}

// Config holds the application configuration
type Config struct {
	Files []FileConfig `json:"files"`
}

const (
	credentialsFile = "credentials.json"
	tokenFile       = "token.json"
	configFile      = "config.json"
)

func loadConfig() (*Config, error) {
	file, err := os.Open(configFile)
	if err != nil {
		return nil, fmt.Errorf("config file not found")
	}
	defer file.Close()

	var config Config
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, fmt.Errorf("error parsing config file: %v", err)
	}
	return &config, nil
}
