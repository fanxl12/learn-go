package main

import (
	"fmt"
	"time"
)

func main() {
	var msg = make(chan int, 0)

	go func(msg chan int) {
		for data := range msg {
			fmt.Println(data)
		}
		fmt.Println("end")
	}(msg)

	msg <- 1
	msg <- 2
	close(msg) // 通知已经发完了 已经关闭的channel可以继续取值，但是不能再放值了
	time.Sleep(time.Second * 10)
}
