package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

func JsonRunner2() {
	Pv("Fake Product 2 Json")
	fakefile2, err := os.Open("./fakeproduct2.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(fakefile2)
	fmt.Println(reflect.TypeOf(fakefile2))

	defer fakefile2.Close()
	byteValuefakefile2, _ := io.ReadAll(fakefile2)
	var fake2 fakeproduct2
	json.Unmarshal(byteValuefakefile2, &fake2)

	fmt.Println(fake2)

	Pv("Pretty Print Json")
	b, err := json.MarshalIndent(fake2, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
