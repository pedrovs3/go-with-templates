package models

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Quantity    int
	Edit        bool
}
