package main

import "fmt"

func main() {
	fmt.Println("Hello")
	// if else, for, switch case, break continue

	f := true
	flag := &f

	if flag == nil {
		fmt.Println("Nil")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// infinite loop
	// for {

	// }

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	arr := []string{"ola", "hi", "asss"}

	for i, value := range arr {
		fmt.Println(i, " - ", value)
	}

	mymap := make(map[string]interface{})
	mymap["name"] = "angad"
	mymap["age"] = 20

	for k, v := range mymap {
		fmt.Println(k, ": ", v)
	}

}
