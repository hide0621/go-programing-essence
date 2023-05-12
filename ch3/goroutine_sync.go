package ch3

import (
	"fmt"
	"sync"
)

func doGoroutineSync1() {
	var wg sync.WaitGroup
	wg.Add(1) // リファレンスカウンターを1つ追加する

	go func() {
		defer wg.Done() // リファレンスカウンターを1つ削除する
		// 何か処理を書く
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}()

	// 何か処理を書く
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// 何か処理を書く
	arr := []string{"foo", "bar", "baz"}
	for index, value := range arr {
		fmt.Println(index, value)
	}

	wg.Wait() // リファレンスカウンターが0になるまで待つ
}

func doGoroutineSync2() {

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		v := 1
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()
}
