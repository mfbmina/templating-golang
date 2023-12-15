package main

import (
	"os"
	"text/template"
)

type Invoice struct {
	Number     string
	Items      []Item
	TotalPrice float64
}

type Item struct {
	Name       string
	Price      float64
	Quantity   int
	TotalPrice float64
}

func main() {
	invoice := Invoice{Number: "1234", Items: []Item{{Name: "Item 1", Price: 12.34, Quantity: 2, TotalPrice: 24.68}, {Name: "Item 2", Price: 56.78, Quantity: 1, TotalPrice: 56.78}}, TotalPrice: 81.46}

	file := "invoice.tmpl"
	tmpl, err := template.New(file).ParseFiles(file)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, invoice)
	if err != nil {
		panic(err)
	}
}
