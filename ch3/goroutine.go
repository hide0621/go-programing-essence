package ch3

import (
	"fmt"
	"time"
)

func sendMessage() {
	message := "h1"

	go func() {
		fmt.Println(message)
	}()
	message = "ho"

	time.Sleep(time.Second)
	println(message)
	time.Sleep(time.Second)
}
