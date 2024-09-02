package main

import (
	"errors"
	"fmt"
	"time"
)

func canGoWrong() (string, error) {
	if time.Now().UnixMilli()%2 == 1 {
		return "Ran Successfully", nil
	}
	return "", errors.New("this is a unidentified error")
}

func main() {
	data, err := canGoWrong()
	if err != nil {
		fmt.Println(err)

		// Exit out the application
		panic("So thats why cant continue further...")
	} else {
		fmt.Println(data)
	}
}
