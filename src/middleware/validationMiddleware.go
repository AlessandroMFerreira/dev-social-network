package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func httpError(w http.ResponseWriter) {
	message := make(map[string]string)
	message["message"] = "Forbidden"
	body, _ := json.Marshal(message)
	http.Error(w, string(body), http.StatusForbidden)
}

func ValidateUser(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI == "/login" || r.RequestURI == "/user/createUser" {
			next.ServeHTTP(w, r)
			return
		}

		tokenString := r.Header.Get("Authorization")
		if len(strings.Split(tokenString, " ")) < 2 {
			httpError(w)
			return
		}

		tokenString = strings.Split(tokenString, " ")[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid token")
			}

			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			httpError(w)
			return
		}

		var expirationDate float64

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			r.Header.Set("id", fmt.Sprint(claims["id"]))

			expirationDate, err = strconv.ParseFloat(fmt.Sprint(claims["exp"]), 64)
			if err != nil {
				httpError(w)
				return
			}

			currentTime := float64(time.Now().Unix())

			if err != nil {
				httpError(w)
				return
			}

			if expirationDate > currentTime {
				next.ServeHTTP(w, r)
			} else {
				httpError(w)
				return
			}
		} else {
			httpError(w)
			return
		}

	})
}
