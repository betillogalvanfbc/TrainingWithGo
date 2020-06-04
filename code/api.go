package main

import (
	"fmt"
	"log"
	"net/http"
)

//BASE URL IS THE BASE ENDPOINT FOR THE STAR WARS API

const BaseURL = "https://swapi.dev/api/"

func getPeople(w http.ResponseWriter, r *http.Request) {

	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Failed to request start warts people")
	}

	fmt.Println(res)
}

func main() {
	http.HandleFunc("/people", getPeople)
	fmt.Println("Serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
