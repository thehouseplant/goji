package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

// Define Configuration structure
type Config struct {
	BaseURL  string `json:"base_url"`
	Email    string `json:"email"`
	APIToken string `json:"api_token"`
}
