package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"reflect"
)

func StructReflect() {
	Pv("Fake Product 2 Json")
	fakefile2, err := os.Open("./fakeproduct2.json")
	if err != nil {
		panic(err)
	}

	defer fakefile2.Close()
	byteValuefakefile2, _ := io.ReadAll(fakefile2)
	var fake2 fakeproduct2
	json.Unmarshal(byteValuefakefile2, &fake2)

	Pv("Reflection")
	ty := reflect.TypeOf(fake2)
	field := ty.Field(0)
	fmt.Println(field)

	fmt.Println("tags", field.Tag.Get(`json`))

	/*
		Pv("REFLECT test ")

		//var products = fake2.Products
		for _, value := range products {
		var reviews = (value.Reviews)
		for _, rat := range reviews {
			var ratings = rat.Rating
			if ratings == 4 {

			}
		}
		}


		t := reflect.TypeOf(fake2)
		fmt.Println(t)
		v := reflect.ValueOf(fake2)

		if t.Kind() != reflect.Struct {
			fmt.Println("Not a struct, check your input")
			return
		}
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			value := v.Field(i)

			fmt.Printf("Field Name: %s, Field Type: %s, Field Value: %v\n",
				field.Name, field.Type, value)

		}

		Pv("Alternative method")
		// Get the reflection value of an object.
		rValue := reflect.ValueOf(fake2)

		// Get the reflection kind of an object.
		rKind := rValue.Kind()
		fmt.Println(rKind)

		// Get the reflection type of an object (two options).
		rType := reflect.TypeOf(fake2)
		//rType := rValue.Type()

		//Alertnative method
		if rType.Kind() == reflect.Struct {
			for i := 0; i < rType.NumField(); i++ {
				fieldValue := rValue.Field(i)
				fmt.Println(fieldValue)
			}

		}
	*/
}
