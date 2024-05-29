package handlers

import "net/http"

func internalServerError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/users", GetUserByID)
	mux.HandleFunc("/search", SearchByCriteria)
	return mux
}
