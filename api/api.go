package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)
// temp define template folder
var temp = template.Must(template.ParseGlob("templates/*.html"))

func StartApi() {
	http.HandleFunc("/", index)
	fmt.Println("Listening on http://localhost:8080/api/v1")
	http.HandleFunc("/api/v1", headers)
	http.HandleFunc("/api/v1/client", clientHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "index.html", nil)
}
