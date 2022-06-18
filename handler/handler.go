package handler

import (
	"net/http"
)

func New() (http.Handler) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("*** Welcome to go-kvstore ***"))
	})
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Everything looks good!"))
	})
	mux.HandleFunc("/db", dbHandler())

	return mux
}

func dbHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.Write([]byte("Get"))
		case http.MethodPost:
			w.Write([]byte("Set"))
		case http.MethodPut:
			w.Write([]byte("Put"))
		case http.MethodDelete:
			w.Write([]byte("Delete"))
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	}
}