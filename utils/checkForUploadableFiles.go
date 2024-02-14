package utils

import (
	"drive-sync/types"
	"fmt"
	"log"
	"os"
	"sync"

	"google.golang.org/api/drive/v3"
)

func CheckForUploadableFiles(srv *drive.Service, DB *types.DBContainer, backupPath string, wg *sync.WaitGroup) {
	readDirRes, readDirErr := os.ReadDir(backupPath)

	if readDirErr != nil {
		log.Fatalf("Could not read directory for uploading. Dir: %v", backupPath)
	}

	for _, file := range readDirRes {
		if file.IsDir() {
			CheckForUploadableFiles(srv, DB, backupPath+"/"+file.Name(), wg)
		}

		info, err := file.Info()

		if err != nil {
			log.Fatalf("Could not get info on file %v. Error: %v", file.Name(), err.Error())
		}

		fmt.Println(info.ModTime())
		// info.ModTime()
		// Check if file entry in DB modified time is less than file modified
		// if so, upload it, else continue
	}
}
