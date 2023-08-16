package controller

import (
	"dev-social-network/src/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user, erro := service.CreateUser(r.Body)
	w.Header().Set("Content-Type", "application/json")

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = erro.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
		return
	}

	body, erro := json.Marshal(user)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func FindAllUsers(w http.ResponseWriter, r *http.Request) {
	var offset string
	var limit string

	if len(r.URL.Query().Get("offset")) <= 0 {
		offset = "0"
	} else {
		offset = r.URL.Query().Get("offset")
	}

	if len(r.URL.Query().Get("limit")) <= 0 {
		limit = "10"
	} else {
		limit = r.URL.Query().Get("limit")
	}

	users, erro := service.FindAll(limit, offset)
	w.Header().Set("Content-Type", "application/json")

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = erro.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
		return
	}
	body, erro := json.Marshal(users)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")

	user, erro := service.FindOne(id)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = erro.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
		return
	}
	body, erro := json.Marshal(user)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user, erro := service.UpdateUser(r.Body)
	w.Header().Set("Content-Type", "application/json")

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = erro.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
		return
	}
	body, erro := json.Marshal(user)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}

func RemoveUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	w.Header().Set("Content-Type", "application/json")

	erro := service.DeleteUser(id)

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = erro.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func Follow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := service.Follow(r.Body, r.Header.Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = err.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func UnFollow(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := service.UnFollow(r.Body, r.Header.Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = err.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func FindFollowing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var offset string
	var limit string

	if len(r.URL.Query().Get("offset")) <= 0 {
		offset = "0"
	} else {
		offset = r.URL.Query().Get("offset")
	}

	if len(r.URL.Query().Get("limit")) <= 0 {
		limit = "10"
	} else {
		limit = r.URL.Query().Get("limit")
	}

	results, err := service.FindFollowing(r.Header.Get("id"), offset, limit)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = err.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
	} else {
		body, err := json.Marshal(results)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}

func FindFollowers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var offset string
	var limit string

	if len(r.URL.Query().Get("offset")) <= 0 {
		offset = "0"
	} else {
		offset = r.URL.Query().Get("offset")
	}

	if len(r.URL.Query().Get("limit")) <= 0 {
		limit = "10"
	} else {
		limit = r.URL.Query().Get("limit")
	}

	results, err := service.FindFollowers(r.Header.Get("id"), offset, limit)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		message := make(map[string]string)
		message["message"] = err.Error()
		body, _ := json.Marshal(message)
		w.Write(body)
	} else {
		body, err := json.Marshal(results)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}
}
