package main

import "fmt"

func main() {
	age1 := 9

	if age1 > 18 {
		fmt.Print("Go and vote vro!")
	} else {
		fmt.Print("Go and eat lollipop for another ", 18-age1, " years!")
	}
}
