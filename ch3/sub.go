package ch3

import "fmt"

func Sub() {
	// list_iota()
	// list_loop()
	// list_slice()
	// list_map()

	// user := User{
	// 	Name: "Bob",
	// 	Age:  18,
	// }

	// list_struct_showName(&user)

	v := Value(1)
	v = v.Add(2)
	fmt.Println(v)
}
