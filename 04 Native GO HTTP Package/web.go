package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// make GET request
	response, error := http.Get("https://reqres.in/api/products")
	if error != nil {
		fmt.Println(error)
	}

	// read response body
	body, error := io.ReadAll(response.Body)
	if error != nil {
		fmt.Println(error)
	}
	// close response body
	response.Body.Close()

	// print response body
	fmt.Println(string(body))
}
