package main

import (
	"fmt"
	"net/http" // the http native package
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprint(w, "<h1>Hello this is the root route</h1>")
	case "/cat":
		fmt.Fprint(w, "<h1>Hello this is the cat route</h1>")
	case "/dog":
		fmt.Fprint(w, "<h1>Hello this is the dog route</h1>")
	}
}

func main() {
	http.HandleFunc("/", handler)     //
	http.ListenAndServe(":3000", nil) // this will start the server in the
}
