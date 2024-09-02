package main

import (
	"fmt"
)

func main() {
	prices := []float64{}
	var qty []float64 = []float64{1, 2, 3, 4} // will create slice and array behind the scenes and shrink expand its size if needed.

	// will create array
	// var qty [4]float64 = [4]float64{1, 2, 3, 4} // also can have size predefined.

	items := []string{}

	// joins 2 slice and makes new one and not actually appends like list/vector.
	items = append(items, "Cello")
	fmt.Println(prices, qty, items)

	fmt.Println(items[0])

	// from index 1 to 2. 3 is excluded
	featuredPrices := qty[1:3]

	fmt.Println(featuredPrices)

	fmt.Println(qty[2:])
	fmt.Println(qty[:3])

	// will get compiletime error if used line 11 instead of 9. with 9 we'll get runtime error.
	// fmt.Println(qty[:5])

	// also editing the slices will edit original array, its reference to window of original array.

	testSlice := qty[:3]

	testSlice[0] = 100

	fmt.Println(qty)

	fmt.Println(len(qty), cap(qty))

	fmt.Println(len(testSlice), cap(testSlice))

	temp := []float64{10.99, 68.2378}

	temp = append(temp, 21.123, 123)

	temp2 := []float64{11230.99, 368.2378}

	// wont work as one list cant be added as single item.
	// temp = append(temp, temp2)

	// we can unpack the list and send individual elements.
	temp = append(temp, temp2...)

	fmt.Println(temp)

	fmt.Println()

	// make a slice and assign to usernames but the behing the scenes array of size 100.
	usernames := make([]string, 100)

	// already initalized to blanks
	fmt.Println(usernames)

}
