package main

import (
	"fmt"
	"net/http" // the http native package
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello world</h1>") // sending a h1 tag
	fmt.Fprintf(w, "This is Rabin")        // sending a string
}

func main() {
	http.HandleFunc("/", handler)     //
	http.ListenAndServe(":3000", nil) // this will start the server in the
}
