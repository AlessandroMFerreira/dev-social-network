package controller

import (
	"api/src/service"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func SavePost(w http.ResponseWriter, r *http.Request) {
	userId := r.Header.Get("id")
	body := r.Body

	post, err := service.SavePost(body, userId)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		message := make(map[string]string)
		message["message"] = err.Error()

		returnedObj, _ := json.Marshal(message)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(returnedObj)

		return
	}

	w.WriteHeader(http.StatusCreated)
	returnedObj, _ := json.Marshal(post)

	w.Write(returnedObj)
}

func UpdatePost(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	body := r.Body

	post, err := service.UpdatePost(body, id)

	w.Header().Set("Content-Type", "application/json")

	if err != nil {
		message := make(map[string]string)
		message["message"] = err.Error()

		returnedObj, _ := json.Marshal(message)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(returnedObj)

		return
	}

	returnedObj, _ := json.Marshal(post)

	w.WriteHeader(http.StatusOK)
	w.Write(returnedObj)
}

func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("id")
	body := r.Body

	w.Header().Set("Content-Type", "application/json")

	err := service.DeletePost(body, id)

	if err != nil {
		message := make(map[string]string)
		message["message"] = err.Error()

		returnedObj, _ := json.Marshal(message)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(returnedObj)

		return
	}

	w.WriteHeader(http.StatusOK)

}

func FindAllPosts(w http.ResponseWriter, r *http.Request) {
	var offset string
	var limit string
	id := r.Header.Get("id")

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

	posts, err := service.FindAllPosts(id, limit, offset)

	if err != nil {
		message := make(map[string]string)
		message["message"] = err.Error()

		returnedObj, _ := json.Marshal(message)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(returnedObj)

		return
	}

	body, _ := json.Marshal(posts)

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func FindPost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["id"]
	userId := r.Header.Get("id")

	post, err := service.FindPost(postId, userId)

	if err != nil {
		message := make(map[string]string)
		message["message"] = err.Error()

		returnedObj, _ := json.Marshal(message)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(returnedObj)

		return
	}

	w.WriteHeader(http.StatusOK)

	returnedObj, _ := json.Marshal(post)

	w.Write(returnedObj)
}

func LikePost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["id"]
	userId := r.Header.Get("id")

	err := service.LikePost(postId, userId)

	if err != nil && strings.Contains(err.Error(), "user already like this post") {
		w.WriteHeader(http.StatusNotModified)
		return
	} else if err != nil && strings.Contains(err.Error(), "there is a dislike for this post") {
		w.WriteHeader(http.StatusNotModified)
		return
	} else if err != nil && strings.Contains(err.Error(), "user already like this post") {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	if err != nil {
		message := make(map[string]string)
		message["message"] = err.Error()

		returnedObj, _ := json.Marshal(message)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(returnedObj)

		return
	}

	w.WriteHeader(http.StatusOK)
}

func DislikePost(w http.ResponseWriter, r *http.Request) {
	postId := mux.Vars(r)["id"]
	userId := r.Header.Get("id")

	err := service.DislikePost(postId, userId)

	if err != nil && strings.Contains(err.Error(), "error on dislike post") {
		w.WriteHeader(http.StatusNotModified)
		return
	} else if err != nil && strings.Contains(err.Error(), "there is a dislike for this post") {
		w.WriteHeader(http.StatusNotModified)
		return
	} else if err != nil && strings.Contains(err.Error(), "user already like this post") {
		w.WriteHeader(http.StatusNotModified)
		return
	}

	if err != nil {
		message := make(map[string]string)
		message["message"] = err.Error()

		returnedObj, _ := json.Marshal(message)

		w.WriteHeader(http.StatusInternalServerError)
		w.Write(returnedObj)

		return
	}

	w.WriteHeader(http.StatusOK)
}
