package main

import (
	"fmt"
	"os"
)

const DB_PATH string = "data.db"

func offloadToFile(something string) {
	os.WriteFile(DB_PATH, []byte(something), 0644)
}

func loadFromFile() string {
	data, error := os.ReadFile(DB_PATH)
	if error != nil {
		return ""
	} else {
		return string(data)
	}
}

func main() {

	var input string

	fmt.Println("Currently stored data: ", loadFromFile())

	fmt.Print("Enter something: ")
	fmt.Scan(&input)
	fmt.Println()
	offloadToFile(input)
	fmt.Println("Saved Successfully.")
}
