package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID                 int
	Title              string
	Description        string
	Price              int
	DiscountPercentage float32
	Rating             float32
	Stock              int
	Images             []string
}

type Response struct {
	Products []Product
	Total    int
	Skip     int
	Limit    int
}

func main() {
	client := &http.Client{}
	resp, err := client.Get("https://dummyjson.com/products")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	response := &Response{}
	json.NewDecoder(resp.Body).Decode(response)

	fmt.Printf("%#v\n", response.Products[0].Title)
}

// Use Read function
// Use io.Copy
// Use io.ReadAll
// Use json.Decoder with a map
// Use json.Decoder with a struct
