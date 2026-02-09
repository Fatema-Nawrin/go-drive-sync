# File Sync

It is a simple command-line tool that simplifies the process of keeping local file synchronized  with Google Drive. 

## Features

- **OAuth2 Authentication**: Secure authentication with Google Drive API
- **Automatic Sync**: Creates new files or updates existing ones

## Prerequisites

1. **Google Cloud Project** with Drive API enabled
2. **OAuth2 Credentials** (credentials.json)

## Setup

### 1. Google Cloud Setup

1. Go to [Google Cloud Console](https://console.cloud.google.com/)
2. Create a new project or select existing one
3. Enable the **Google Drive API**
4. Create OAuth2 credentials:
   - Go to "Credentials" → "Create Credentials" → "OAuth client ID"
   - Select "Desktop application"
   - Download the JSON file and rename it to `credentials.json`

### 2. Configuration

Create a `config.json` file following config.json.example



##  Run the application:

```bash
go build
./file-sync
# Sync only a specific file
./file-sync -file "file1.pdf"
```

**First run only**: 
   - Open the provided URL in your browser
   - Grant permission to access Google Drive
   - Copy the authorization code and paste it in the terminal

Subsequent runs will use the cached token automatically
