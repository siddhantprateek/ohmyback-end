package handler

import (
	"fmt"
	"net/http"
)

type Data struct {
	ID        int      `json:"id,omitempty"`
	UserID    int      `json:"userId,omitempty"`
	UserName  string   `json:"userName,omitempty"`
	Timestamp int64    `json:"timestamp,omitempty"`
	TxnType   string   `json:"txnType,omitempty"`
	Amount    string   `json:"amount,omitempty"`
	Location  Location `json:"location,omitempty"`
	IP        string   `json:"ip,omitempty"`
}

type JsonData struct {
	Page       int    `json:"page,omitempty"`
	PerPage    int    `json:"per_page,omitempty"`
	Total      int    `json:"total,omitempty"`
	TotalPages int    `json:"total_pages,omitempty"`
	Data       []Data `json:"data,omitempty"`
}
type Location struct {
	ID      int    `json:"id,omitempty"`
	Address string `json:"address,omitempty"`
	City    string `json:"city,omitempty"`
	ZipCode int    `json:"zipCode,omitempty"`
}

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
