package api

import (
	"api/src/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func UserApi(router *mux.Router) *mux.Router {
	router.HandleFunc("/user/createUser", controller.CreateUser).Methods(http.MethodPost)

	router.HandleFunc("/user/findAllUsers", controller.FindAllUsers).Methods(http.MethodGet)

	router.HandleFunc("/user/findUser/{id}", controller.FindUser).Methods(http.MethodGet)

	router.HandleFunc("/user/updateUser", controller.UpdateUser).Methods(http.MethodPost)

	router.HandleFunc("/user/removeUser/{id}", controller.RemoveUser).Methods(http.MethodDelete)

	router.HandleFunc("/user/follow", controller.Follow).Methods(http.MethodPost)

	router.HandleFunc("/user/unfollow", controller.UnFollow).Methods(http.MethodPost)

	router.HandleFunc("/user/findFollowing", controller.FindFollowing).Methods(http.MethodGet)

	router.HandleFunc("/user/findFollowers", controller.FindFollowers).Methods(http.MethodGet)

	return router
}
