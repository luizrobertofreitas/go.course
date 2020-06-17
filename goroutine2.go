package main

import (
	"fmt"
	"sync"
)

func main() {
	// wait groups
	wg := &sync.WaitGroup{}

	wg.Add(2)

	go func() {
		fmt.Println("Lambda")
		wg.Done()
	}()

	go func() {
		fmt.Println("Lambda 2")
		wg.Done()
	}()

	fmt.Println("Main")
	wg.Wait()
	fmt.Println("Finished")
}
