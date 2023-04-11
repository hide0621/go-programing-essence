package main

import (
	"fmt"
)

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

func list_iota(a int) {
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

	// fmt.Println(Apple)
	// fmt.Println(Orange)
	// fmt.Println(Banana)

	var fruit Fruit = Apple
	fmt.Println(fruit)

	// fruit = Elephant
	// fmt.Println(fruit)

}

func list_loop(b int) {

outerLoop:
	for i := 0; i < 5; i++ {
		fmt.Printf("Outer loop iteration %d\n", i)

		for j := 0; j < 5; j++ {
			fmt.Printf("  Inner loop iteration %d\n", j)

			// Inner loopの3回目でouterLoopから脱出
			if i == 2 && j == 2 {
				break outerLoop
			}
		}
	}

	fmt.Println("Done")
}

func main() {

	list_iota(1)

	list_loop(1)

}
