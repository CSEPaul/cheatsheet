package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
)

func JsonRunner() {
	Pv("JSON")

	file, err := os.Open("./fakeproduct.json")
	if err != nil {
		panic(err)
	}
	fmt.Println(file)
	fmt.Println(reflect.TypeOf(file))

	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

	Pv("Unmarshal with a manual method and structure is known")
	var users Users
	json.Unmarshal(byteValue, &users)
	/*
		if err != nil {
			log.Printf("error decoding response: %v", err)
			if e, ok := err.(*json.SyntaxError); ok {
				log.Printf("Syntax error at byte offset %d", e.Offset)
			}
			log.Printf("response:%q", file)
			fmt.Println(err)
		}*/
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}

	Pv("Unmarshal when struct is known using a struct")
	fmt.Println(users.Users[0].Name)
	for key, values := range users.Users {
		fmt.Println(key, values)
		if values.Name == "Elliot" {
			fmt.Println("Elliot Found")
		}
	}

	Pv("Unmarshal when struct is unknown")
	var dat map[string]interface{}
	if err := json.Unmarshal(byteValue, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

}
