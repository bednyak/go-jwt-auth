package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/bednyak/go-react-jwt-auth/pkg/errors"
	"github.com/golang-jwt/jwt"
	"net/http"
	"os"
)

func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {
			var err errors.Error
			err = errors.SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(os.Getenv("JWT_SECRET_KEY"))

		token, err := jwt.Parse(r.Header["Authorization"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err errors.Error
			err = errors.SetError(err, "Your authorization token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if claims["role"] == "admin" {
				r.Header.Set("Role", "admin")
				handler.ServeHTTP(w, r)
				return

			} else if claims["role"] == "user" {
				r.Header.Set("Role", "user")
				handler.ServeHTTP(w, r)
				return

			}
		}
		var reserr errors.Error
		reserr = errors.SetError(reserr, "Not Authorized.")
		json.NewEncoder(w).Encode(err)
	}
}
