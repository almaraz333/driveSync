package utils

import (
	"drive-sync/types"
	"time"

	"google.golang.org/api/drive/v3"
)

func CheckIfShouldDownload(file *drive.File, DB *types.DBContainer) bool {
	dbFile, foundFile := DB.Items[file.Id]

	if !foundFile {
		return true
	}

	DBFileT, _ := time.Parse("2006-01-02T15:04:05.999Z", dbFile.LastUpdated)

	driveFileT, _ := time.Parse("2006-01-02T15:04:05.999Z", file.ModifiedTime)

	if driveFileT.Unix() > DBFileT.Unix() {
		return true
	}

	return false
}
