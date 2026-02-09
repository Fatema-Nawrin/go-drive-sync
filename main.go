package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	specificFile = flag.String("file", "", "Sync only a specific file by flag")
)

func main() {
	flag.Parse()

	// Load configuration
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Validate we have files to sync
	if len(config.Files) == 0 {
		log.Fatalf("No files configured in config.json")
	}

	// Get Google Drive service
	srv, err := getDriveService()
	if err != nil {
		log.Fatalf("Error getting Drive service: %v", err)
	}

	// Determine which files to sync
	filesToSync := config.Files
	if *specificFile != "" {
		// User wants to sync a specific file
		found := false
		for _, fc := range config.Files {
			if fc.DriveFileName == *specificFile {
				filesToSync = []FileConfig{fc}
				found = true
				break
			}
		}
		if !found {
			log.Fatalf("File '%s' not found in config.json", *specificFile)
		}
	}

	// Sync files
	successCount := 0
	failCount := 0

	for _, fc := range filesToSync {
		// Validate local file exists
		if _, err := os.Stat(fc.LocalFile); os.IsNotExist(err) {
			failCount++
			continue
		}

		// Sync the file
		if err := syncFile(srv, &fc); err != nil {
			fmt.Printf("Failed: %v\n", err)
			failCount++
		} else {
			fmt.Printf("Synced successfully\n")
			successCount++
		}
	}

	if failCount > 0 {
		os.Exit(1)
	}
}
