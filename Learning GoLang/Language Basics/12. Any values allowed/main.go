package main

import "fmt"

// dangerous
func acceptAnything(value interface{}) {
	// accepts anything although not useful.
	fmt.Println(value)
}

func main() {
	acceptAnything(129213)
	acceptAnything("Jugal")
	acceptAnything("23.235")
}
