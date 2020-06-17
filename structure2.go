package main

import (
	"errors"
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	day := "Fry"

	switch day {
	case "Fry":
		fmt.Println("Sextou")
	case "Mon", "Tue", "Wed", "Thu":
		fmt.Println("Time to work")
	default:
		fmt.Println("Weekend")
	}

	err := errors.New("alksdjflkajsdfklj")
	fmt.Println(err)
}
