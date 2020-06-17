package main

import "fmt"

func main() {
	var c chan int
	c = make(chan int) // or c := make(chan int)

	// send

	go func() {
		c <- 1
	}()

	// sniff

	val := <-c
	fmt.Println(val)

	go func() {
		c <- 2
	}()

	val = <-c
	fmt.Println(val)

	fmt.Println(c)
}
