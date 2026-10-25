package main

import (
	"fmt"
	"time"
)

var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 1
	for {
		//我怎么去做到，应该此处，等待另一个goroutine来通知我
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}

func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPORSTUVWXYZ"
	for {
		//我怎么去做到，应该此处，等待另一个goroutine来通知我
		if i >= len(str) {
			return
		}
		<-letter
		fmt.Printf(str[i : i+2])
		i += 2
		number <- true
	}
}

func main() {
	/*
		使用两个goroutine交替打印序列，一个goroutine打印数字，另外一个goroutine打印字母，
		最终效果如下:
		12AB34C056EF78GH910IJ1112KL1314MN15160P1718QR1920ST2122UV2324WX2526YZ2728
	*/
	go printNum()
	go printLetter()
	number <- true
	time.Sleep(time.Second * 100)
}
