package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
}

func getUser() (*User, error) {
	return &User{ID: 1, FirstName: "John", LastName: "Doe"}, nil
}

func getUserProfile(userID int) (*User, error) {
	if userID != 1 {
		return nil, errors.New("User profile not found")
	}
	return &User{ID: 1, FirstName: "John", LastName: "Doe"}, nil
}

func userController() {
	user, err := getUser()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	profile, err := getUserProfile(user.ID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("User:", user)
	fmt.Println("Profile:", profile)
}

// go has a built-in error interface type

// create a new error
// var err error = errors.New("error message")
