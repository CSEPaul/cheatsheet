package main

import (
	"encoding/json"
	"fmt"
)

func Ps() {
	fmt.Println(" ")
}

func Pv(value string) {
	fmt.Println("----", value, "-----")
}

func main() {
	Pv("strings and numbers")
	fmt.Println("test print strings")
	fmt.Println("Numbers", 123)

	Ps()
	Pv("Arrays")
	// Array of length 10
	fmt.Println("Array of length 10")
	var a [10]int
	fmt.Println(a)

	// Array of length 4 and with values
	fmt.Println("Array of length 10 and with values")
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)

	// Array of length 10 and with values
	var b = [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	fmt.Println(b)
	fmt.Println("first element", b[1])
	fmt.Println("length of array", len(b))
	fmt.Println("-----")
	fmt.Println("Type of: ")
	fmt.Printf("%T\n", b)
	fmt.Println("-----")
	fmt.Println("last element of array", b[len(b)-1])

	// Maps
	var c = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	d := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", c)
	fmt.Println("length of map", len(c))
	fmt.Printf("b\t%v\n", d)

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
