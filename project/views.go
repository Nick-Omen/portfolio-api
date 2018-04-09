package project

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

func SetRoutes(router *mux.Router) {
	router.HandleFunc("/project/", getProjects).Methods("GET")
	router.HandleFunc("/project/{id}/", getProject).Methods("GET")
	router.HandleFunc("/project/", createProject).Methods("POST")
	router.HandleFunc("/project/{id}/", updateProject).Methods("PUT")
	router.HandleFunc("/project/{id}/", deleteProject).Methods("DELETE")
}

func getProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := m.GetAllProjects()
	if err == nil {
		json.NewEncoder(w).Encode(projects)
	}
}

func getProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if idStr, ok := vars["id"]; ok {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			project, err := m.GetProjectById(id)
			if err == nil {
				json.NewEncoder(w).Encode(project)
			} else {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
			}
		}
	}
}

func createProject(w http.ResponseWriter, r *http.Request) {
	project := &Project{}
	_ = json.NewDecoder(r.Body).Decode(project)
	project, err := m.CreateProject(project)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(project)
	}
}

func updateProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if idStr, ok := vars["id"]; ok {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			project, err := m.GetProjectById(id)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
			} else {
				json.NewDecoder(r.Body).Decode(project)
				m.UpdateProject(project)
				json.NewEncoder(w).Encode(project)
			}
		}
	}
}

func deleteProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if idStr, ok := vars["id"]; ok {
		id, err := strconv.Atoi(idStr)
		if err == nil {
			project, err := m.GetProjectById(id)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				json.NewEncoder(w).Encode(err.Error())
			} else {
				m.DeleteProject(project)
			}
		}
	}
}
