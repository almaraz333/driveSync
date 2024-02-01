package utils

import (
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/api/drive/v3"
)

func GetFoldersFolders(parentFolder *drive.File, srv *drive.Service, initDir string, waitGroup *sync.WaitGroup) {
	folderQueryString := fmt.Sprintf("mimeType = 'application/vnd.google-apps.folder' and 'me' in owners and '%v' in parents", parentFolder.Id)

	res, err := srv.Files.List().Q(folderQueryString).Do()

	if err != nil {
		log.Fatalf("Could not get folder's foldres. Parent Folder ID: %v", parentFolder.Id)
	}

	for _, val := range res.Files {
		currDir := initDir + "/" + val.Name
		os.MkdirAll(currDir, 0755)
		GetFoldersFolders(val, srv, currDir, waitGroup)
	}

	fileQueryString := fmt.Sprintf("mimeType != 'application/vnd.google-apps.folder' and 'me' in owners and '%v' in parents", parentFolder.Id)
	fileRes, fileErr := srv.Files.List().Q(fileQueryString).Do()

	if fileErr != nil {
		log.Fatalf("Could not get files from folder with ID %v", parentFolder.Id)
	}

	for _, val := range fileRes.Files {
		fmt.Println("FILE:", val.Name, val.MimeType)
		if !val.Trashed {
			waitGroup.Add(1)
			go DownloadFile(srv, val, initDir, waitGroup)
		}
	}
}
