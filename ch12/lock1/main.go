package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var total int32
var wg sync.WaitGroup

func add() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		// 原子包 增加操作原子化
		atomic.AddInt32(&total, 1)
	}
}

func sub() {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
		atomic.AddInt32(&total, -1)
	}
}

func main() {
	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(total)
}
