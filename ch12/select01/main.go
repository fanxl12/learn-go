package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{}) //channel 是多线程安全的, channel要初始化

func g1() {
	time.Sleep(time.Second)
	done <- struct{}{}
}

func g2() {
	time.Sleep(time.Second * 2)
	done <- struct{}{}
}

func main() {
	//select 类似于 switch case语句， 但是select的功能和我们操作Linux里面提供的io的select、poll、epoll
	//select 主要作用于多个channel

	//现在有个需求，我们现在有两个goroutine都在执行，但是呢我在主的goroutine中，当某一个执行完成以后，这个时候我会立马知道
	go g1()
	go g2()

	<-done
	fmt.Println("done")
}
