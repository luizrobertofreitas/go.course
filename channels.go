package main

import "fmt"

func publish(n int, c chan int) {
	go func() {
		c <- n
	}()
}

func main() {
	c := make(chan int)
	// <name> chan <datatype>

	// send in a goroutine
	publish(1, c)
	publish(2, c)
	publish(3, c)
	publish(4, c)
	publish(5, c)

	// snif
	val1 := <-c
	val2 := <-c
	val3 := <-c
	val4 := <-c
	val5 := <-c
	fmt.Println(val1, val2, val3, val4, val5)

	fmt.Println(c)
}
