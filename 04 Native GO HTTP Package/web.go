package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Read the cookie named "myCookie"
	cookie, err := r.Cookie("Apple")
	if err != nil {
		// Cookie not found
		fmt.Fprintf(w, "Cookie not found")
		return
	}

	// Cookie found, print its value
	fmt.Fprintf(w, "Cookie Value: %s", cookie.Value)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{
		Name:  "Apple",
		Value: "Its a fruit",
	}
	http.SetCookie(w, &cookie)
	w.Write([]byte("Cookie set successfully!"))
}

func main() {
	http.HandleFunc("/get-cookie", handler) // In this route we can access the cookie
	http.HandleFunc("/", setCookie)         // the root route sets up a cookie

	http.ListenAndServe(":8080", nil)
}
