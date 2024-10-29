package friend

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func addFriendsRequestHandler(srv FriendService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var friend Friend
		err := json.NewDecoder(r.Body).Decode(&friend)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&ResponseFriendRespons{Success: false})
			return
		}

		success, err := srv.addFriendRequest(friend)

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&ResponseFriendRespons{Success: success, Message: err})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(&ResponseFriendRespons{Success: success})
	}
}

func FriendHandler(router *mux.Router, srv FriendService) {
	router.HandleFunc("/friend-request", addFriendsRequestHandler(srv)).Methods("POST")
}
