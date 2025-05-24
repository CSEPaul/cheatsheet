package main

import "fmt"

func mapper() {
	Ps()
	Pv("MAPS")
	var c = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	d := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", c)
	fmt.Println("length of map", len(c))
	fmt.Printf("b\t%v\n", d)

}
