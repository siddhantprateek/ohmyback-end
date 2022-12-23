package main

import (
	"log"
	"net/http"

	handler "github.com/siddhantprateek/ohmyback-end/go_jwt_auth/handler"
	JWT "github.com/siddhantprateek/ohmyback-end/go_jwt_auth/jwt"
)

func main() {
	// endpoint: /api
	http.Handle("/api", JWT.ValidateToken(handler.HomeRoute))
	http.HandleFunc("/jwt", JWT.GetJWT)
	// listen to port 8000
	log.Printf("Listening... to port http://localhost:8000/api")
	http.ListenAndServe(":8000", nil)

}
