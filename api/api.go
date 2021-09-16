package api

import (
	"fmt"
	"log"
	"net/http"
)

func headers(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	fmt.Fprintf(w, "Hello ")
	for _, v := range q["name"] {
		fmt.Fprintf(w, "%s ", v)
	}
	fmt.Fprintf(w, "\n")
}

func StartApi(){
	fmt.Println("Listening on http://localhost:8080/api/v1")
	http.HandleFunc("/api/v1", headers)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
