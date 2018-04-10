package storage

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"strings"
)

func Connect(name string) *gorm.DB {
	db, err := gorm.Open("sqlite3", strings.Join([]string{name, ".db"}, ""))
	if err != nil {
		log.Panic(err)
	}
	return db
}