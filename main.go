package main

import (
	"LetsGooDocs/cmd"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading .env file")
	}
	
	cmd.Execute()
}
