package service

import (
	"api/src/model"
	"api/src/repository"
	"api/src/utils"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func ValidateAndReturnJwt(w http.ResponseWriter, r *http.Request) {
	var userLogin model.UserLogin
	var retrivedUserLogin model.UserLogin

	data, erro := io.ReadAll(r.Body)

	w.Header().Set("Content-Type", "application/json")

	if erro != nil {
		w.WriteHeader(http.StatusUnauthorized)
		message := make(map[string]string)
		message["message"] = "Invalid email or password"
		body, _ := json.Marshal(message)
		w.Write(body)
		return
	}

	erro = json.Unmarshal(data, &userLogin)

	if erro != nil {
		w.WriteHeader(http.StatusUnauthorized)
		message := make(map[string]string)
		message["message"] = "Invalid email or password"
		body, _ := json.Marshal(message)
		w.Write(body)
		return
	}

	dbConnection, erro := utils.OpenConnection()

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer dbConnection.Close()

	retrivedUserLogin, erro = repository.LogIn(dbConnection, userLogin.Email)

	if erro != nil {
		w.WriteHeader(http.StatusUnauthorized)
		message := make(map[string]string)
		message["message"] = "Invalid email or password"
		body, _ := json.Marshal(message)
		w.Write(body)
		return
	}

	if erro = bcrypt.CompareHashAndPassword([]byte(retrivedUserLogin.Password), []byte(userLogin.Password)); erro != nil {
		w.WriteHeader(http.StatusUnauthorized)
		message := make(map[string]string)
		message["message"] = "Invalid email or password"
		body, _ := json.Marshal(message)
		w.Write(body)
		return
	}

	expirationDate := time.Now().Add(time.Hour)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  retrivedUserLogin.Id,
		"exp": expirationDate.Unix(),
	})

	tokenString, erro := token.SignedString([]byte(os.Getenv("SECRET")))

	if erro != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	cookie := http.Cookie{Name: "Authorization", Value: tokenString, Expires: expirationDate}

	http.SetCookie(w, &cookie)
}
