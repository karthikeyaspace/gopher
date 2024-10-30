package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	user := User{
		ID:   1,
		Name: "Karthikeya",
		Age:  19,
	}

	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error converting data to json", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, string(data))
}

func httpServer() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/user", userHandler)
	http.ListenAndServe(":8080", nil)
}
