package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	numbers = processNumber(&numbers, returnAppropriateFunc(2))
	fmt.Println(numbers)
	numbers = processNumber(&numbers, returnAppropriateFunc(3))
	fmt.Println(numbers)
	// passing anonymous func as paramter
	numbers = processNumber(&numbers, func(i int) int { return i * 10 })
	fmt.Println(numbers)
	numbers = processNumber(&numbers, returnAppropriateFunc(11))
	fmt.Println(numbers)
	numbers = processNumber(&numbers, returnAppropriateFunc(-1))
	fmt.Println(numbers)
}

// taking func as parameter
func processNumber(numbers *[]int, functionToBeCalled func(int) int) []int {
	res := []int{}
	for _, val := range *numbers {
		res = append(res, functionToBeCalled(val))
	}
	return res
}

func doubleNumber(number int) int {
	return number * 2
}

func tripleNumber(number int) int {
	return number * 3
}

// returning a func.
func returnAppropriateFunc(times int) func(int) int {
	if times == 2 {
		return doubleNumber
	}
	if times == 3 {
		return tripleNumber
	} else {
		// anonymous func
		return func(number int) int { return number * times }
	}
}
