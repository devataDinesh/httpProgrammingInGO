package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type College struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	fmt.Println("Starting Server!!!")

	http.HandleFunc("/Dinesh", welcomeHandler)

	http.HandleFunc("/GetInfo", getInfoHandler)

	http.HandleFunc("/PostData", postInfoHandler)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to Server!")
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	college := College{
		Name: "Amrita University",
		City: "Kollam",
	}

	err := json.NewEncoder(w).Encode(&college)
	if err != nil {
		fmt.Println("Some error in Fetching the Info", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func postInfoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Application/json")

	college := College{}

	err := json.NewDecoder(r.Body).Decode(&college)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Data obtained from body is ", college)
}
