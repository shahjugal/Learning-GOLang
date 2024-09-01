package main

import "fmt"

func add[T int | float64 | float32 | string](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(add(1, 2), add("Jugal ", "Shah"), add(69.9, -0.9))
}
