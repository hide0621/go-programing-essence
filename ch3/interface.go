package ch3

import "fmt"

func PrintDetail(v interface{}) {

	switch t := v.(type) {
	case int, int32, int64:
		fmt.Println("int/int32/int64 型:", t)
	case string:
		fmt.Println("string型:", t)
	default:
		fmt.Println("知らない型")
	}
}
