package benefit

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func getBenefitsHandler(srv BenefitService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		category := vars["category"]
		u, err := srv.GetBenefits(category)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(&BenefitRespons{Benefits: u})
	}
}

func BenefitHandler(router *mux.Router, srv BenefitService) {
	router.HandleFunc("/benefits/{category}", getBenefitsHandler(srv)).Methods("GET")
}
