package types

import "sync"

type DBItem struct {
	Id             string `json:"id"`
	LastUpdated    string `json:"lastUpdated"`
	ParentFolderId string `json:"parentFolderId"`
	FilePath       string `json:"filePath"`
}

type DBContainer struct {
	Mu    sync.Mutex
	Items map[string]DBItem
}

func (db *DBContainer) lockDB() {
	db.Mu.Lock()
}

func (db *DBContainer) unlockDB() {
	db.Mu.Unlock()
}
