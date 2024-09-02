package main

import (
	"fmt"
	"time"
)

func greet(phrase string, doneChan chan bool) {
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func slowGreet(phrase string, doneChan chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	doneChan <- true
}

func main() {

	// runs serially
	/*
		greet("Nice to meet you!")
		greet("How are you?")
		slowGreet("How ... are ... you ...?")
		greet("I hope you're liking the course!")
	*/

	// make a channel which uses bool for communication in this case whether its done or not.

	dones := make([]chan bool, 2)
	dones[0] = make(chan bool)
	go slowGreet("How ... are ... you ...?", dones[0])
	dones[1] = make(chan bool)
	go greet("How are you?", dones[1])

	// it will wait till any data comes out of the both channel.
	for _, done := range dones {
		<-done
	}

}
