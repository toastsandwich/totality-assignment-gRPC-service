package handlers

import (
	"net/http"

	"github.com/WistleBlowers/totality-assignment-RESTful-version/service"
	"github.com/WistleBlowers/totality-assignment-RESTful-version/utils"
)

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	usr, err := service.GetUserByID(id)
	if err != nil {
		internalServerError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	utils.CreateAndWriteJSON(w, usr)
}

func SearchByCriteria(w http.ResponseWriter, r *http.Request) {
	criteria := r.URL.Query().Get("criteria")
	value := r.URL.Query().Get("value")
	usrs, err := service.SearchByCriteria(criteria, value)
	if err != nil {
		internalServerError(w, err)
	}
	w.Header().Set("Content-Type", "application/json")
	utils.CreateAndWriteJSON(w, usrs)
}
