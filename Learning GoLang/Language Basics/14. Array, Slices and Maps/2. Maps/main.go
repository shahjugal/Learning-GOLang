package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google": "https://google.com",
		"Amazon": "https://aws.com",
	}

	fmt.Println(websites)

	fmt.Println(websites["Google"])

	websites["Microsoft"] = "https://microsoft.com"

	fmt.Println(websites)

	delete(websites, "Google")

	fmt.Println(websites)

	// make a hashmap and assign to courses but the capacity will be 100 so we wont make new map and rehash. less collisions.
	courses := make(map[string]bool, 100)

	courses["React"] = true
	courses["Angular"] = true
	courses["JSS"] = true

	// already initalized to blanks
	fmt.Println(courses)

	// for getting key value
	for key, val := range courses {
		fmt.Println(key, val)
	}

	// for key only
	for pair := range courses {
		fmt.Println(pair)
	}

}
