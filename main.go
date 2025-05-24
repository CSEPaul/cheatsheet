package main

import (
	"fmt"
)

func main() {

	fmt.Println("test print strings")
	fmt.Println("Numbers", 123)

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
	fmt.Println(b[1])
	fmt.Println(b[-1])

}
