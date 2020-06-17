package main

import "fmt"

// func main() {
// 	var m1 int // type explicit
// 	m1 = 1

// 	var (
// 		m2 = 2 // implicit datatype
// 		m3 = 3
// 	)

// 	var m4 int32 // default initialized to 0
// 	var m5 int64

//   // To make the following operation, it is needed to cast
// 	fmt.Println(int64(m4) + m5)

// 	// declaring and initializing
// 	m6 := 3

// 	fmt.Println("Hello World")
// }

// Strings
// func main() {
// 	m1 := "my name amooora"
// 	m2 := "name"
// 	fmt.Println(m1)
// 	fmt.Println(strings.Contains(m1, m2))
// 	fmt.Println(strings.Replace(m1, "a", "b", -1))
// 	fmt.Println(strings.Split(m1, " "))
// }

// Arrays
// func todo() {
// 	// var arr []int
// 	arr := []int{1, 2, 3, 4}
// 	arr2 := []string{"hi", "my", "name"}
// 	arr2 = append(arr2, "is", "angad", "!")
// 	fmt.Println(arr, arr2)
// }

// func main() {
// 	todo()
// }

// Pointers

// func swap(m1, m2 *int) {
// 	var temp int
// 	temp = *m2
// 	*m2 = *m1
// 	*m1 = temp
// }

// func main() {
// 	m2 := 2
// 	m1 := 1
// 	ptr := &m2
// 	fmt.Println(ptr)  // prints the address
// 	fmt.Println(*ptr) // using address to get value
// 	swap(&m1, &m2)
// 	fmt.Println(m1, m2)
// }

// Structs
type Car struct {
	Name    string
	Age     int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("Driving...")
}

func (c Car) GetName() string {
	return c.Name
}

func main() {
	c := Car{
		Name:    "chevy",
		Age:     1,
		ModelNo: 2,
	}
	c.Print()
	c.Drive()
	fmt.Println(c.GetName())
}
