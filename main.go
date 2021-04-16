package main

import (
	"log"
	"os"
)

func main() {
	// first thing we need is a file name
	fName := "data.csv"
	// creating a file
	file, err := os.Create(fName)
	if err != nil {
		// Fatalf basically prints the message and exits the program
		log.Fatalf("File could not be created, error : %q", err)
		return
	}
	defer file.Close()
}
