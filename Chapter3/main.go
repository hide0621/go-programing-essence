package main

import "fmt"

// const (
// 	Apple = iota
// 	Orange
// 	Banana
// )

// const (
// 	Apple = iota + iota
// 	Orange
// 	Banana
// )

// const (
// 	Apple = iota + iota
// 	Orange
// 	Banana = iota + 3
// )

type Fruit int
type Animal int

const (
	Apple Fruit = iota
	Orange
	Banana
)

const (
	Monckey Animal = iota
	Elephant
	Pig
)

func main() {
	// fmt.Println(Apple)
	// fmt.Println(Orange)
	// fmt.Println(Banana)

	var fruit Fruit = Apple
	fmt.Println(fruit)

	// fruit = Elephant
	// fmt.Println(fruit)
}
