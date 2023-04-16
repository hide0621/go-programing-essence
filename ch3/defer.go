package ch3

import "fmt"

// 複数個deferがあれば逆順で実行される
func SampleDefer() {
	defer fmt.Println(4)
	defer fmt.Println(5)
	defer fmt.Println(6)
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
}

// deferは無名関数でも実行できる
func doSomething() {
	var n int = 1
	defer func() {
		fmt.Println(n)
	}()
	n = 2
}
