package utils

import (
	"drive-sync/types"

	"google.golang.org/api/drive/v3"
)

func AddFileToDB(driveFile *drive.File, filePath string, DB *types.DBContainer) {

	var parentFolderId string

	if len(driveFile.Parents) > 0 {
		parentFolderId = driveFile.Parents[0]
	}

	test := types.DBItem{
		FilePath:       filePath,
		ParentFolderId: parentFolderId,
		LastUpdated:    driveFile.ModifiedTime,
		Id:             driveFile.Id,
	}

	DB.Mu.Lock()
	defer DB.Mu.Unlock()

	DB.Items[test.Id] = test
}
