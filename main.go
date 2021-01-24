package main

import (
	"log"

	clone "clone-repos/clone"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	clone.StartCloneProcess()
}
