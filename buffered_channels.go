package main

import "fmt"

type Car struct {
	Model string
}

func main() {
	var c chan *Car
	c = make(chan *Car, 3) // or c := make(chan int)

	// send

	go func() {
		c <- &Car{"1"}
		c <- &Car{"2"}
		c <- &Car{"3"}
		c <- &Car{"4"}
		close(c)
	}()

	// sniff

	for car := range c {
		fmt.Println(car.Model)
	}

	fmt.Println(c)
}
