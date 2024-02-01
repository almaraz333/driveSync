# Google Drive File Downloader

This Go project utilizes the Google Drive API and Go routines to download your Google Drive files and replicate the directory structure locally. The script will download the files to the same location where the script is located.

## Features

- Download Google Drive files locally while preserving the directory structure.
- Utilizes Go routines for concurrent downloads, improving performance.
- Recursively traverses through directories to maintain the structure.

## Getting Started

To get started with this project, follow the steps below:

### Prerequisites

- Install Go on your system. You can download it from the official [Go website](https://golang.org/dl/).

### Installation

1. Clone the repository:

```
git clone https://github.com/almaraz333/driveSync.git
```

2. Navigate to the project directory:

```
cd driveSync
```

3. Install the project dependencies:

```
go mod download
```

### Configuration

Before running the script, you need to set up Google Drive API credentials. Follow these steps:

1. Visit the [Google Cloud Console](https://console.cloud.google.com/).

2. Create a new project or select an existing one.

3. Enable the Google Drive API for your project.

4. Create credentials:
   - Go to "Credentials" under the "APIs & Services" section.
   - Click "Create Credentials" and select "OAuth client ID".
   - Choose "Desktop app" as the application type.
   - Note down the generated client ID and client secret.

5. Copy the `credentials.json` file from the Google Cloud Console into the project's root directory.

### Usage

To run the script, use the following command:

```
go run main.go
```

## Contributing

Contributions are welcome! If you encounter any bugs or have suggestions for improvements, please create an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- This project relies on the [Google Drive API](https://developers.google.com/drive).
