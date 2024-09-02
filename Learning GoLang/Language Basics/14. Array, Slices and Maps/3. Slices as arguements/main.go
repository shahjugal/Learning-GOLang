package main

import "fmt"

func main() {
	fmt.Println(sumup([]int{1, 2, 3, 4, 5}))
	fmt.Println(sumAll(1, 2, 3, 4, 5))

	fmt.Println(sumAll([]int{1, 2, 3, 4, 5}...))

	/*
		... can be used to spread out the list to individuals
	*/

}

func sumup(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// this is also called variadic functions.
func sumAll(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
