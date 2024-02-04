package main

import (
	m "alura_store/models"
	u "alura_store/utils"
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	u.Check(err)
}

func index(w http.ResponseWriter, _ *http.Request) {
	products := []m.Product{
		{
			Name:        "Apple MacBook Pro 17",
			Description: "Silver",
			Price:       2999.99,
			Quantity:    120,
			Edit:        true,
		},
		{
			Name:        "Magic Mouse 2",
			Description: "Black",
			Price:       99.99,
			Quantity:    3,
		},
		{
			Name:        "Google Pixel Phone",
			Description: "Gray",
			Price:       799,
			Quantity:    43,
		},
	}

	err := temp.ExecuteTemplate(w, "Index", products)
	u.Check(err)
}
