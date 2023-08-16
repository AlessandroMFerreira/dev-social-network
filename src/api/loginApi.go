package api

import (
	"dev-social-network/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func LoginApi(router *mux.Router) *mux.Router {
	router.HandleFunc("/login", controller.Login).Methods(http.MethodPost)

	return router
}
