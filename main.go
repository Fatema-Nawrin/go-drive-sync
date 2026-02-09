package main

import (
	"log"
)

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	srv, err := getDriveService()
	if err != nil {
		log.Fatalf("Error creating Drive service: %v", err)
	}

	if err := syncFile(srv, config); err != nil {
		log.Fatalf("Error syncing file: %v", err)
	}

	log.Println("File synced successfully!")
}
