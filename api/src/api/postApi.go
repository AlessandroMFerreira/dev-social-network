package api

import (
	"api/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func PostApi(router *mux.Router) *mux.Router {
	router.HandleFunc("/post/save", controller.SavePost).Methods(http.MethodPost)

	router.HandleFunc("/post/update", controller.UpdatePost).Methods(http.MethodPost)

	router.HandleFunc("/post/deletePost", controller.DeletePost).Methods(http.MethodDelete)

	router.HandleFunc("/post/findAllPosts", controller.FindAllPosts).Methods(http.MethodGet)

	router.HandleFunc("/post/findPost/{id}", controller.FindPost).Methods(http.MethodGet)

	return router
}
