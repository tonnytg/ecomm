package api

import (
	"fmt"
	"log"
	"net/http"
)

func StartApi() {
	fmt.Println("Listening on http://localhost:8080/api/v1")
	http.HandleFunc("/api/v1", headers)
	http.HandleFunc("/api/v1/client", clientHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
