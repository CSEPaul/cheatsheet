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

	Pv("MAPS - appending")
	data := map[string]string{"new": "1"}
	data["new2"] = "2"
	fmt.Println(data)

	Pv("MAPS - removing")
	delete(data, "new2")
	fmt.Println(data)

	Pv("MAPS - Get keys")
	if x, found := d["Oslo"]; found {
		fmt.Println(x)
		fmt.Println(found)
	}

	Pv("MAPS - Get All keys and values")
	for key, value := range d {
		fmt.Println(key)
		fmt.Println(value)
	}

}
