package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/tidwall/gjson"
)

func JsonFilter() {
	Pv("JsonFilter with gjson")

	fake1, err := os.Open("./fakeproduct.json")
	if err != nil {
		panic(err)
	}
	defer fake1.Close()
	byteValuefakefile1, _ := io.ReadAll(fake1)

	// Make an instance of fakeproduct1
	var fake11 fakeproduct1
	json.Unmarshal(byteValuefakefile1, &fake11)

	//Marshal to prepare convert to string
	fmt.Println(reflect.TypeOf(fake11))
	jsonmarshal, err := json.Marshal(fake11)
	if err != nil {
		panic(err)
	}

	// Convert to string
	jsonstring := string(jsonmarshal)
	//fmt.Println(jsonstring)

	//Gjson apply filter
	allnames := gjson.Get(jsonstring, "users.#.name")
	fmt.Println(allnames)

	socialmediaTwitter := gjson.Get(jsonstring, "users.#.social.twitter")
	fmt.Println(socialmediaTwitter)

}
