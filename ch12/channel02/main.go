package main

import (
	"fmt"
	"time"
)

func prducer(out chan<- int) {
	for i := range 10 {
		out <- i * i
	}
}

func consumer(in <-chan int) {
	for data := range in {
		fmt.Printf("num-%d\r\n", data)
	}
}

func main() {
	//默认情况下，channel是双向的
	// 但是，我们经常一个channel作为参数进行传递，希望对方是单向使用

	//var ch1 chan int //双向channel
	//var ch2 chan<-float64 //单向channel，只能写入float64的数据
	//var ch3<-chan int//单向的，只能读取

	//c := make(chan int,3)
	//var sender chan<- int = c // send-only
	//var reader <-chan int = c // recv-only
	//
	//sender <- 1
	//<-reader

	c := make(chan int)
	go prducer(c)

	go consumer(c)

	time.Sleep(time.Second * 10)

}
