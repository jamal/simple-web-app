package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func Hello(w http.ResponseWriter, r *http.Request) {
	me := Person{"Jamal", "Fanaian"}
	response, err := json.Marshal(me)
	if err != nil {
		fmt.Fprintf(w, "Error: "+err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, string(response))
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}
