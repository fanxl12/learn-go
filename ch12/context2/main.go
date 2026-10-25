package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

// 我们新需求，可以主动退出监控
func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
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
	// context包提供了三种函数，withCancel, withTimeout, withValue
	//如果你的goroutine， 函数中，如果希望被控制， 超时、传值，但是我不希望影响我原来的接口信息的时候，函数参数中第一个参数就尽量的要加上一Context
	ctx1, _ := context.WithCancel(context.Background())
	ctx2, _ := context.WithTimeout(ctx1, time.Second*6)
	go cpuInfo(ctx2)

	wg.Wait()
	fmt.Println("监控完成")
}
