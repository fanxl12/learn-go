package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 我们新需求，可以主动退出监控
// 共享变量
var stop bool

func cpuInfo() {
	defer wg.Done()
	for {
		if stop {
			break
		}
		time.Sleep(time.Second)
		fmt.Println("cpu info")
	}
}

func main() {
	// 渐进式的方式
	wg.Add(1)
	go cpuInfo()
	time.Sleep(time.Second * 6)
	stop = true
	wg.Wait()
	fmt.Println("监控完成")
}
