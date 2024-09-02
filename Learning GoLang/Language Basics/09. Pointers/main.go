package main

import "fmt"

func main() {
	// Adding 20 to age2
	withoutPointers()

	fmt.Println("Now using pointers")

	withPointers()

	var val *int

	fmt.Println("Trying to dereference the nullptr(nil for golang): ", val)

	fmt.Println("Value of dereferenced nil", *val)

}

func withoutPointers() {
	age := 32
	fmt.Println("age: ", age)

	age2 := age
	fmt.Println("age2: ", age2)

	age2 += 20
	fmt.Println("age2: ", age2)
	fmt.Println("age: ", age)
}

func withPointers() {
	age := 32
	fmt.Println("age: ", age)

	age2 := &age
	fmt.Println("age2: ", *age2)

	*age2 += 20
	fmt.Println("age2: ", *age2)
	fmt.Println("age: ", age)
}
