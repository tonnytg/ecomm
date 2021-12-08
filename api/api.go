package api

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// tpl define template folder
var tpl = template.Must(template.ParseGlob("templates/*.html"))

func StartApi() {
	http.HandleFunc("/", index)
	fmt.Println("Listening on http://localhost:8080/api/v1")
	http.HandleFunc("/api/v1", headers)
	http.HandleFunc("/api/v1/client", clientHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type Produto struct {
	Name string
	Price float64
	Actions []string
}

func index(w http.ResponseWriter, r *http.Request) {

	produtos := []Produto{
			{Name: "Produto 1", Price: 10.00, Actions: []string{"Comprar", "Vender"}},
			{Name: "Produto 2", Price: 11.00, Actions: []string{"Comprar", "Vender"}},
		}

	tpl.ExecuteTemplate(w, "index.html", produtos)
}
