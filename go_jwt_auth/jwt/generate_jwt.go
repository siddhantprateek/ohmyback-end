package jwt

import (
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var SECRET = []byte("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")
var API_KEY = "<provide-api-key>(any value)"

// JWT token generator
func GenerateJWT() (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	// for the payload
	claims := token.Claims.(jwt.MapClaims)

	// allocating expiration time
	claims["exp"] = time.Now().Add(time.Hour).Unix()

	tokenStr, err := token.SignedString(SECRET)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return tokenStr, nil

}

// validate token
func ValidateToken(next func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(t *jwt.Token) (interface{}, error) {
				_, ok := t.Method.(*jwt.SigningMethodHMAC)
				if !ok {
					w.WriteHeader(http.StatusUnauthorized)
					w.Write([]byte("not authorized"))
				}
				return SECRET, nil
			})

			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("not authorized: " + err.Error()))
			}

			if token.Valid {
				next(w, r)
			}
		} else {
			// if there is no token in the header
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("not authorized"))
		}
	})
}

// @desp: Get JWT token
func GetJWT(w http.ResponseWriter, r *http.Request) {
	if r.Header["Access"] != nil {
		if r.Header["Access"][0] != API_KEY {
			return
		} else {
			token, err := GenerateJWT()
			if err != nil {
				return
			}
			fmt.Fprint(w, token)
		}
	}
}
