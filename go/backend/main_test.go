package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"testing"
)

func TestMain(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	main()
}
