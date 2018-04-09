package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nick_omen_api/storage"
	"nick_omen_api/project"
	"github.com/jinzhu/gorm"
)

func shareConnection(db *gorm.DB)  {
	project.SetDatabase(db)
}

func shareRouter(router *mux.Router)  {
	project.SetRoutes(router)
}

func main () {
	router := mux.NewRouter()
	db := storage.Connect()
	shareConnection(db)
	shareRouter(router)

	log.Fatal(http.ListenAndServe(":9000", router))
}
