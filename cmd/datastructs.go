package cmd

import "time"

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
type fakeproduct1 struct {
	Users []struct {
		Name   string `json:"name"`
		Type   string `json:"type"`
		Age    int    `json:"age"`
		Social struct {
			Facebook string `json:"facebook"`
			Twitter  string `json:"twitter"`
		} `json:"social"`
	} `json:"users"`
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
