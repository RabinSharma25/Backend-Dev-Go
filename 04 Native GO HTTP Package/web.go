// Building a simple API using the http package

package main

import (
	"encoding/json"
	"fmt"
	"net/http" // the http native package
)

type weath struct {
	Place     string `json:"place"`
	Longi     int    `json:"longi"`
	Lati      int    `json:"lati"`
	Temp      int    `json:"temp"`
	Condition string `json:"condition"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the homepage endpoint hit")
}

func weather(w http.ResponseWriter, r *http.Request) {
	var w1 = weath{
		Place:     "Chisopani",
		Longi:     23,
		Lati:      24,
		Temp:      35,
		Condition: "Humid",
	}
	json.NewEncoder(w).Encode(w1) // converion of struct to json
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/weather", weather)
	http.ListenAndServe(":3000", nil) // this will start the server in the the port 3000
}
