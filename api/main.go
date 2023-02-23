package main

import (
	"api/src/api"
	"api/src/middleware"
	"api/src/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	utils.Load()

	routes := mux.NewRouter()
	routes = api.UserApi(routes)
	routes = api.LoginApi(routes)
	routes = api.PostApi(routes)
	routes.Use(middleware.ValidateUser)

	fmt.Printf("Server listening on port %d", utils.ApiPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", utils.ApiPort), routes))
}
