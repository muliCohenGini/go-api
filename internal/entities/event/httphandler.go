package event

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getEventsHandler(srv EventService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := srv.GetEvents()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(&EventRespons{Events: u})
	}
}

func EventHandler(router *mux.Router, srv EventService) {
	router.HandleFunc("/events", getEventsHandler(srv)).Methods("GET")
}
