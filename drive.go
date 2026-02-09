package main

import (
	"context"
	"fmt"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v2"
	"google.golang.org/api/option"
)

func getDriveService() (*drive.Service, error) {
	ctx := context.Background()
	b, err := os.ReadFile(credentialsFile)
	if err != nil {
		return nil, fmt.Errorf("Error reading credentials file")
	}
	config, err := google.ConfigFromJSON(b, drive.DriveFileScope)
	if err != nil {
		return nil, fmt.Errorf("unable to parse credentials: %v", err)
	}
	client := getClient(config)
	srv, err := drive.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return nil, fmt.Errorf("unable to create Drive service: %v", err)
	}
	return srv, nil
}

func findFileInFolder(srv *drive.Service, fileName, folderID string) (string, error) {
	query := fmt.Sprintf("title='%s' and '%s' in parents", fileName, folderID)
	fmt.Printf("Searching for file with query: %s\n", query)
	fileList, err := srv.Files.List().Q(query).Fields("items(id, title)").Do()
	if err != nil {
		return "", fmt.Errorf("error searching for file: %v", err)
	}

	fmt.Printf("Found %d files\n", len(fileList.Items))
	if len(fileList.Items) > 0 {
		return fileList.Items[0].Id, nil
	}

	return "", nil
}

func syncFile(srv *drive.Service, config *FileConfig) error {
	fileID, err := findFileInFolder(srv, config.DriveFileName, config.DriveFolderID)
	if err != nil {
		return err
	}
	file, err := os.Open(config.LocalFile)
	if err != nil {
		return fmt.Errorf("error opening local file: %v", err)
	}
	defer file.Close()

	if fileID != "" {
		// Update existing file
		fmt.Println("File exists, updating...")
		_, err = srv.Files.Update(fileID, &drive.File{}).Media(file).Do()
		if err != nil {
			return fmt.Errorf("error updating file: %v", err)
		}
	} else {
		// Create new file
		fmt.Println("File doesn't exist, creating...")
		fileMetadata := &drive.File{
			Title:   config.DriveFileName,
			Parents: []*drive.ParentReference{{Id: config.DriveFolderID}},
		}
		_, err = srv.Files.Insert(fileMetadata).Media(file).Do()
		if err != nil {
			return fmt.Errorf("error creating file: %v", err)
		}
	}

	return nil
}
