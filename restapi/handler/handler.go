package handler

import (
	"fmt"
	"net/http"
)

func HelloHandler(resp http.ResponseWriter, req *http.Request) {

	fmt.Fprint(resp, "Hello there!")
	resp.WriteHeader(http.StatusOK)
}

// path: http://localhost:8080/query?name=John%20Doe
/*
/  desp: query string is part of the URL which can be sued toadd some data
/  to the request for the resource.
*/
// > curl http://localhost:8000/query?name=siddhant
func QueryParamHandler(resp http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")

	if name == "" {
		name = "guest"
	}

	fmt.Fprintf(resp, "Hello %s!", name)
}
