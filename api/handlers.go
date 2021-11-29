package api

import (
	"ecomm/database"
	client "ecomm/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Host: %s\n", r.Host)
}

func clientHandler(w http.ResponseWriter, r *http.Request) {

	// GET it is to obtain information about the client
	if r.Method == "GET" {
		// get query
		q := r.URL.Query()
		// if you don't have a name, return error bad request
		if q["name"] == nil {
			status := http.StatusBadRequest
			http.Error(w, http.StatusText(status), status)
			return
		}
		// loop through the query each name founded
		for _, v := range q["name"] {
			fmt.Fprintf(w, "GET: Hello %s from the client handler\n", v)
		}
	}

	// POST it is to create a new client
	if r.Method == "POST" {

		// get data from json
		var formattedBody client.Client
		err := json.NewDecoder(r.Body).Decode(&formattedBody)
		if err != nil {
			status := http.StatusBadRequest
			http.Error(w, http.StatusText(status), status)
			return
		}

		// get the body of the request and register a new client
		c := client.NewClient(formattedBody.Username, formattedBody.Password, formattedBody.Email)
		if c.Check() {
			database.DB.Create(&c)

			status := http.StatusCreated
			http.Error(w, http.StatusText(status), status)
			return
		}

		status := http.StatusBadRequest
		http.Error(w, "must be send username, password and email", status)
		return
	}
}
