package main

import (
	"github.com/gorilla/mux"
	"nick_omen_api/storage"
	"log"
	"net/http"
	"github.com/jinzhu/gorm"
	"nick_omen_api/project"
	"nick_omen_api/tag"
	"os"
)

func shareConnection(db *gorm.DB)  {
	project.SetDatabase(db)
	tag.SetDatabase(db)
}

func shareRouter(router *mux.Router)  {
	project.SetRoutes(router)
	tag.SetRoutes(router)
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			next.ServeHTTP(w, r)
		} else {
			authenticated := true  // TODO add authentication check
			if authenticated {
				next.ServeHTTP(w, r)
			} else {
				w.WriteHeader(http.StatusForbidden)
			}
		}
	})
}

func startProduction() {
	router := mux.NewRouter()
	router.Use(authMiddleware)
	db := storage.Connect("api")
	db.LogMode("true" == os.Getenv("DEV"))
	shareConnection(db)
	shareRouter(router)

	log.Fatal(http.ListenAndServe(":9000", router))
}

func startTest() {
	router := mux.NewRouter()
	db := storage.Connect("test")
	shareConnection(db)
	shareRouter(router)

	log.Fatal(http.ListenAndServe(":9002", router))
}

func main () {
	if "true" == os.Getenv("TEST") {
		startTest()
	} else {
		startProduction()
	}
}
