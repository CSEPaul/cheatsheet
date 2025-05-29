package cmd

import (
	"encoding/json"
	"fmt"
)

func Strut() {

	Pv("STRUCTS")

	// structs
	type people struct {
		age  string
		name string
	}

	var p people
	p.age = "34"
	p.name = "meme"
	fmt.Println(p)

	// Marshal the data
	peopledata, err := json.Marshal(p)
	if err != nil {
		panic(err)
	}
	fmt.Println("Marshalled data")
	fmt.Println(peopledata)

	fmt.Println(" ")
	fmt.Println(" ")
	// Marshal example 2

	type User struct {
		Name        string
		Age         int
		Active      bool
		lastLoginAt string
	}
	u, err := json.Marshal(User{Name: "Bob", Age: 10, Active: true, lastLoginAt: "today"})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(u)) // {"Name":"Bob","Age":10,"Active":true}

}
