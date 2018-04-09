package storage

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"github.com/jinzhu/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "api.db")
	if err != nil {
		log.Panic(err)
	}
	return db
}