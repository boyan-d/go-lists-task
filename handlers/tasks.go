package handlers

import (
	"encoding/json"
	"net/http"
)

type Task struct {
	ListName string `json:"list"`
	Name     string `json:"name"`
	Done     string `json:"done"`
}

var TasksList []Task

func Create(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	TasksList = append(TasksList, t)

	w.WriteHeader(http.StatusOK)
}

func GetList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	err := json.NewEncoder(w).Encode(TasksList)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func Update(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i := range TasksList {
		if TasksList[i].Name == t.Name && TasksList[i].ListName == t.ListName {

			TasksList[i] = t

			w.WriteHeader(http.StatusOK)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
}
