// Serving files statically from a folder

package main

import (
	"net/http" // the http native package
)

func handler(w http.ResponseWriter, r *http.Request) { // this function will help us serve a serve an entire static folder
	p := "." + r.URL.Path
	if p == "./" {
		p = "./static/index.html" // here the folder we are serving is named "static"
	}
	http.ServeFile(w, r, p)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil) // this will start the server in the the port 3000
}

//////////////////////// The above functionality can be also achieved from the code block given below ////////////////////

/*

package main

import (
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(":3000", nil)
}

*/
