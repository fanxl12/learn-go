package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 我们新需求，可以主动退出监控
func cpuInfo(stop chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stop:
			fmt.Println("监控退出，停止...")
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("cpu info")
		}
	}
}

func main() {
	// 渐进式的方式
	wg.Add(1)
	var stop = make(chan struct{})
	go cpuInfo(stop)
	time.Sleep(time.Second * 6)
	stop <- struct{}{}
	wg.Wait()
	fmt.Println("监控完成")
}
