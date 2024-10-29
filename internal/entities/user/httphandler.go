package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muliCohenGini/go-api/internal/utils"
)

func getUserHandler(srv UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		u, err := srv.GetUser(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Println(err)
			return
		}

		json.NewEncoder(w).Encode(u)
	}
}

func getUsersHandler(srv UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := srv.GetUsers()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(u)
	}
}

func setUserOnboardingDetails(srv UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		response := utils.Response{}
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println("Error decoding request body:", err)
			response.Success = false
			json.NewEncoder(w).Encode(response)
			return
		}

		err = srv.SetUserOnboardingDetails(user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println("Error setting user onboarding details:", err)
			response.Success = false
			response.Message = err.Error()
			json.NewEncoder(w).Encode(response)
			return
		}

		response.Success = true
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}

func checkNicknameHandler(srv UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := utils.Response{}
		vars := mux.Vars(r)
		nickname := vars["nickname"]
		valid, err := srv.GetNicknameValidation(nickname)
		response.Success = valid

		if err != nil {
			response.Message = err.Error()
			switch response.Message {
			case "nickname not valid":
				w.WriteHeader(http.StatusBadRequest)
			case "user already exists":
				w.WriteHeader(http.StatusConflict)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		if valid {
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}
}

func getInterestsHandler(srv UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u, err := srv.GetInterests()
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(u)
	}
}

func updateInterestsHandler(srv UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var interest Interest
		vars := mux.Vars(r)
		id := vars["id"]
		err := json.NewDecoder(r.Body).Decode(&interest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		u, err := srv.UpdateInterests(interest, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		json.NewEncoder(w).Encode(u)
	}
}

func getIdentityNumberValidationHandler(srv UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := utils.Response{}
		var interest Interest
		vars := mux.Vars(r)
		id := vars["id"]
		err := json.NewDecoder(r.Body).Decode(&interest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}
		valid, err := srv.GetIdentityNumberValidation(id)
		response.Success = valid
		if err != nil {
			response.Message = err.Error()
			switch response.Message {
			case "identity number not valid":
				w.WriteHeader(http.StatusBadRequest)
			case "identity number already exists":
				w.WriteHeader(http.StatusConflict)
			default:
				w.WriteHeader(http.StatusInternalServerError)
			}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(response)
	}
}

func UserHandler(router *mux.Router, srv UserService) {
	router.HandleFunc("/users", getUsersHandler(srv)).Methods("GET")
	router.HandleFunc("/users/{id}", getUserHandler(srv)).Methods("GET")
	router.HandleFunc("/user-nickname-validation/{nickname}", checkNicknameHandler(srv)).Methods("GET")
	router.HandleFunc("/users", setUserOnboardingDetails(srv)).Methods("POST")
	router.HandleFunc("/user-interests", getInterestsHandler(srv)).Methods("GET")
	router.HandleFunc("/user-interests/{id}", updateInterestsHandler(srv)).Methods("PUT")
	router.HandleFunc("/user-identity-number/{id}", getIdentityNumberValidationHandler(srv)).Methods("GET")
}
