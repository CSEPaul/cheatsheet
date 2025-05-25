package cmd

import "fmt"

func Arry() {
	Pv("ARRAYS")
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

}
