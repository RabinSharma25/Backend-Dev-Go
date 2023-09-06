// Serving a single static html file

package main

import (
	"net/http" // the http native package
)

func handler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handler)     //
	http.ListenAndServe(":3000", nil) // this will start the server in the the port 3000
}
