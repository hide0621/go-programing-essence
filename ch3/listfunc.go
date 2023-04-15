package ch3

import "fmt"

type Fruit int
type Animal int

type User struct {
	Name string
	Age  int
}

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

func list_slice() {

	a := []int{1, 2, 3}
	fmt.Println(a)

	a = append(a, 4, 5, 6)
	fmt.Println(a)

	fmt.Printf("aの長さは%d\n", len(a))

}

func list_map() {

	m := map[string]int{
		"John": 21,
		"Bob":  18,
		"Mark": 33,
	}
	fmt.Println(m)
}

func list_struct_showName(u *User) {

	fmt.Println(u.Name)
}
