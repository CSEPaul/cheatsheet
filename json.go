package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Products []struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct {
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func jsonRunner() {
	Pv("JSON")

	file, err := os.ReadFile("json.go")
	if err != nil {
		panic(err)
	}
	fmt.Println(file)

	/*
		Pv("Unmarshal when struct is known")
		var p Products
		data := json.Unmarshal(file, &p)
		fmt.Println(data)
	*/

	Pv("Unmarshal when struct is unknown")
	var dat map[string]interface{}
	if err := json.Unmarshal(file, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

}
