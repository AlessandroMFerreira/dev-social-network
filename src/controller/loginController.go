package controller

import (
	"dev-social-network/src/service"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	service.ValidateAndReturnJwt(w, r)
}
