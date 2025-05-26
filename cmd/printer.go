package cmd

import "fmt"

func Ps() {
	fmt.Println(" ")
}

func Pv(value string) {
	fmt.Println("         ")
	fmt.Println("--------------------")
	fmt.Println("----", value, "-----")
	fmt.Println("--------------------")
	fmt.Println("         ")
}
