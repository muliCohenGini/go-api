package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muliCohenGini/go-api/internal/middleware"
	"github.com/muliCohenGini/go-api/internal/router"
)

func main() {
	r := mux.NewRouter()
	router.Routes(r)
	log.Fatal(http.ListenAndServe(":8000", middleware.JsonContentTypeMiddleware(r)))
}
