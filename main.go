package main

import (
	d "alura_store/database"
	m "alura_store/models"
	u "alura_store/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"net/http"
	"os"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8000", nil)
	u.Check(err)
}

func index(w http.ResponseWriter, _ *http.Request) {
	db := d.ConnectWithDatabase()

	selectAllProducts, errorQuery := db.Query("SELECT * FROM tbl_products")
	u.Check(errorQuery)

	p := m.Product{}
	var products []m.Product

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64
		var edit bool

		errorQuery = selectAllProducts.Scan(&id, &name, &description, &price, &quantity, &edit)
		u.Check(errorQuery)

		p.Name = name
		p.Price = price
		p.Edit = edit
		p.Description = description
		p.Quantity = quantity

		products = append(products, p)
	}
	err := temp.ExecuteTemplate(w, "Index", products)
	u.Check(err)
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
	}(db)
}
