package main

import (
	"fmt"
	"time"
)

func g1(c chan struct{}) {
	time.Sleep(time.Second)
	c <- struct{}{}
}

func g2(c chan struct{}) {
	time.Sleep(time.Second * 2)
	c <- struct{}{}
}

func main() {
	//select 类似于 switch case语句， 但是select的功能和我们操作Linux里面提供的io的select、poll、epoll
	//select 主要作用于多个channel

	//现在有个需求，我们现在有两个goroutine都在执行，但是呢我在主的goroutine中，当某一个执行完成以后，这个时候我会立马知道
	c1 := make(chan struct{})
	go g1(c1)
	c2 := make(chan struct{})
	go g2(c2)

	//我要监控多个channel， 任何一个channel返回都知道
	//1.某一个分支就绪了就执行该分支 2.如果两个都就绪了，先执行哪个，随机的，目的是什么，防止饥饿
	// 应用场景 针对g1 和 g2的执行，设置超时时间，如果超时还未完成就结束
	timer := time.NewTimer(time.Second)
	for {
		select {
		case <-c1:
			fmt.Println("g1 done")
		case <-c2:
			fmt.Println("g2 done")
		case <-timer.C:
			fmt.Println("timer expired")
			return
		}
	}

}
