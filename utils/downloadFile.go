package utils

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"google.golang.org/api/drive/v3"
)

func DownloadFile(srv *drive.Service, file *drive.File, dir string, waitGroup *sync.WaitGroup) {
	var httpRes *http.Response
	var err error

	defer waitGroup.Done()

	exportMimeTypes := map[string]bool{
		"application/vnd.google-apps.presentation": true,
		"application/vnd.google-apps.document":     true,
		"application/vnd.google-apps.spreadsheet":  true,
	}

	if _, ok := exportMimeTypes[file.MimeType]; ok {
		httpRes, err = srv.Files.Export(file.Id, "application/pdf").Download()
	} else {
		httpRes, err = srv.Files.Get(file.Id).Download()
	}
	defer httpRes.Body.Close()

	if err != nil {
		log.Fatalf("Could not download file from google drive: %v", err.Error())
	}

	newFilePath := filepath.FromSlash(dir + "/" + file.Name)
	out, osCreateErr := os.Create(newFilePath)
	defer out.Close()

	if osCreateErr != nil {
		log.Fatalf("Could not create file on the OS: %v", osCreateErr.Error())
	}

	_, copyErr := io.Copy(out, httpRes.Body)

	if copyErr != nil {
		log.Fatalf("Could not copy file to OS with file path: %v: %v", dir, copyErr.Error())
	}

}
