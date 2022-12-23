package handler

import (
	"fmt"
	"net/http"
)

func HomeRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint is running perfectly")
}
