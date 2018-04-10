package tag

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/tag/", getTags).Methods("GET")
	router.HandleFunc("/tag/{id}/", getTag).Methods("GET")
	router.HandleFunc("/tag/", createTag).Methods("POST")
	router.HandleFunc("/tag/{id}/", updateTag).Methods("PUT")
	router.HandleFunc("/tag/{id}/", deleteTag).Methods("DELETE")
}

func getTags(w http.ResponseWriter, r *http.Request) {
	tags, err := m.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(tags)
	}
}

func getTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	
	if id, ok := vars["id"]; ok {
		id, err := strconv.Atoi(id)
		if err == nil {
			tag, err := m.GetOne(id)
			if err == nil {
				json.NewEncoder(w).Encode(tag)
			} else {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func createTag(w http.ResponseWriter, r *http.Request) {
	tag := &Tag{}
	_ = json.NewDecoder(r.Body).Decode(&tag)
	tag, err := m.Create(tag)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(tag)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	}
}

func updateTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		id, err := strconv.Atoi(id)
		if err == nil {
			tag, err := m.GetOne(id)
			if err == nil {
				json.NewDecoder(r.Body).Decode(tag)
				tag, err = m.Update(tag)
				if err == nil {
					json.NewEncoder(w).Encode(tag)
				} else {
					w.WriteHeader(http.StatusBadRequest)
					json.NewEncoder(w).Encode(err.Error())
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
			}
		}
	}
}

func deleteTag(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		id, err := strconv.Atoi(id)
		if err == nil {
			tag, err := m.GetOne(id)
			if err == nil {
				deleted := m.Delete(tag)
				if !deleted {
					w.WriteHeader(http.StatusBadRequest)
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
			}
		}
	}
}