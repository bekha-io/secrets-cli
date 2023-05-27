package db

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


var Database *gorm.DB

const (
	databasePath = "./secrets.db"
	databaseEncryptionKey = "fe817e7a-a018-4129-b351-4b9686233a75"
)

func SetupDatabase() {
	var err error
	Database, err = gorm.Open(sqlite.Open(fmt.Sprintf("file:%s?_key=%s&_pragma_key=pass&_pragma_cipher_page_size=4096", 
		databasePath, databaseEncryptionKey)))
	if err != nil {
		log.Fatalf("Cannot open database: %v", err)
	}

	Database.AutoMigrate(&Group{}, &Secret{})
}