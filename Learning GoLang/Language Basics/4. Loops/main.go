package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// continue condition
		if i%2 == 1 {
			continue
		}
		fmt.Println(i + 1)
	}

	i := 0
	for {
		// will run infinitely.
		fmt.Println(i + 1)

		i++

		// break condition
		if i >= 1000 {
			break
		}

	}
	j := 0
	for j < 100 {
		fmt.Println("j is still smaller so running. for ref j is ", j)
		j++
	}
}
