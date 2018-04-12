package project

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	"github.com/asaskevich/govalidator"
	"nick_omen_api/server"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/project/", getProjects).Methods("GET")
	router.HandleFunc("/project/{id}/", getProject).Methods("GET")
	router.HandleFunc("/project/", createProject).Methods("POST")
	router.HandleFunc("/project/{id}/", updateProject).Methods("PUT")
	router.HandleFunc("/project/{id}/", deleteProject).Methods("DELETE")
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := M.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		json.NewEncoder(w).Encode(projects)
	}
}

func getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		id, err := strconv.Atoi(id)
		if err == nil {
			project, err := M.GetOne(id)
			if err == nil {
				json.NewEncoder(w).Encode(project)
			} else {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
			}
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	}
}

func createProject(w http.ResponseWriter, r *http.Request) {
	project := &Project{}
	_ = json.NewDecoder(r.Body).Decode(project)
	valid, err := govalidator.ValidateStruct(project)

	if !valid {
		server.ResponseValidationError(err.Error(), w)
	} else {
		project, err = M.Create(project)
		if err == nil {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(project)
		} else {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err.Error())
		}
	}
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		id, err := strconv.Atoi(id)
		if err == nil {
			project, err := M.GetOne(id)
			if err == nil {
				json.NewDecoder(r.Body).Decode(project)
				valid, err := govalidator.ValidateStruct(project)

				if !valid {
					server.ResponseValidationError(err.Error(), w)
				} else {
					project, err = M.Update(project)
					if err == nil {
						json.NewEncoder(w).Encode(project)
					} else {
						w.WriteHeader(http.StatusBadRequest)
						json.NewEncoder(w).Encode(err.Error())
					}
				}
			} else {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
			}
		}
	}
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if id, ok := vars["id"]; ok {
		id, err := strconv.Atoi(id)
		if err == nil {
			project, err := M.GetOne(id)
			if err == nil {
				deleted := M.Delete(project)
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
