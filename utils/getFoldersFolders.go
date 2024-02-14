package utils

import (
	"drive-sync/types"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/api/drive/v3"
)

func GetFoldersFolders(parentFolder *drive.File, srv *drive.Service, initDir string, waitGroup *sync.WaitGroup, DB *types.DBContainer) {
	folderQueryString := fmt.Sprintf("mimeType = 'application/vnd.google-apps.folder' and 'me' in owners and '%v' in parents", parentFolder.Id)

	res, err := srv.Files.List().Q(folderQueryString).Do()

	if err != nil {
		log.Fatalf("Could not get folder's foldres. Parent Folder ID: %v", parentFolder.Id)
	}

	_, initDirErr := os.Stat(initDir)

	if initDirErr != nil {
		os.Mkdir(initDir, 0755)
	}

	for _, val := range res.Files {
		currDir := initDir + "/" + val.Name
		_, err := os.Stat(currDir)
		if err != nil {
			os.MkdirAll(currDir, 0755)
		}
		GetFoldersFolders(val, srv, currDir, waitGroup, DB)
	}

	fileQueryString := fmt.Sprintf("mimeType != 'application/vnd.google-apps.folder' and 'me' in owners and '%v' in parents", parentFolder.Id)
	fileRes, fileErr := srv.Files.List().Fields("files(name, id, mimeType, parents, trashed, modifiedTime)").Q(fileQueryString).Do()

	if fileErr != nil {
		log.Fatalf("Could not get files from folder with ID %v, Error: %v", parentFolder.Id, fileErr.Error())
	}

	for _, val := range fileRes.Files {
		if !val.Trashed && CheckIfShouldDownload(val, DB) {
			fmt.Println("FILE:", val.Name, val.MimeType)
			waitGroup.Add(1)
			go DownloadFile(srv, val, initDir, waitGroup, DB)
		}
	}
}
