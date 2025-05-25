package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"time"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

type fakeproduct2 struct {
	Products []struct {
		ID                 int      `json:"id"`
		Title              string   `json:"title"`
		Description        string   `json:"description"`
		Category           string   `json:"category"`
		Price              float64  `json:"price"`
		DiscountPercentage float64  `json:"discountPercentage"`
		Rating             float64  `json:"rating"`
		Stock              int      `json:"stock"`
		Tags               []string `json:"tags"`
		Brand              string   `json:"brand,omitempty"`
		Sku                string   `json:"sku"`
		Weight             int      `json:"weight"`
		Dimensions         struct {
			Width  float64 `json:"width"`
			Height float64 `json:"height"`
			Depth  float64 `json:"depth"`
		} `json:"dimensions"`
		WarrantyInformation string `json:"warrantyInformation"`
		ShippingInformation string `json:"shippingInformation"`
		AvailabilityStatus  string `json:"availabilityStatus"`
		Reviews             []struct {
			Rating        int       `json:"rating"`
			Comment       string    `json:"comment"`
			Date          time.Time `json:"date"`
			ReviewerName  string    `json:"reviewerName"`
			ReviewerEmail string    `json:"reviewerEmail"`
		} `json:"reviews"`
		ReturnPolicy         string `json:"returnPolicy"`
		MinimumOrderQuantity int    `json:"minimumOrderQuantity"`
		Meta                 struct {
			CreatedAt time.Time `json:"createdAt"`
			UpdatedAt time.Time `json:"updatedAt"`
			Barcode   string    `json:"barcode"`
			QrCode    string    `json:"qrCode"`
		} `json:"meta"`
		Images    []string `json:"images"`
		Thumbnail string   `json:"thumbnail"`
	} `json:"products"`
	Total int `json:"total"`
	Skip  int `json:"skip"`
	Limit int `json:"limit"`
}

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

}
