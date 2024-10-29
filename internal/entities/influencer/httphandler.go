package influencer

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getInfluencersHandler(srv InfluencerService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := srv.GetInfluencers()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(&InfluencerRespons{Influencers: u})
	}
}

func InfluencerHandler(router *mux.Router, srv InfluencerService) {
	router.HandleFunc("/influencers", getInfluencersHandler(srv)).Methods("GET")
}
