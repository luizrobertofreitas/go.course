package main

import (
	"fmt"
	"time"
)

func heavy() {
	for {
		time.Sleep(time.Second * 1)
		fmt.Println("Heavy")
	}
}

func superHeavy() {
	for {
		time.Sleep(time.Second * 2)
		fmt.Println("Super Heavy")
	}
}

func main() {
	// heavy and superHeavy executes concurrently
	go heavy()
	go superHeavy()
	fmt.Println("Main finished")
	select {} // Waits/Holds the execution like infinite loop for {}
}
