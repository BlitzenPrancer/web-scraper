package main

import "os"

func main() {
	// first thing we need is a file name
	file := "data.csv"
	// creating a file
	file, err := os.Create(file)
}
