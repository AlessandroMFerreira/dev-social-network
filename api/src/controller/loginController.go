package controller

import (
	"api/src/service"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	service.ValidateAndReturnJwt(w, r)
}
