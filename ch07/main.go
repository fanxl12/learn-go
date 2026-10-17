package main

import (
	"fmt"
	"strconv"
)

type MyInt int // 自定义类型

func (mi MyInt) string() string {
	return strconv.Itoa(int(mi))
}

func main() {
	// type 关键字
	/**
	1. 定义结构体
	2. 定义接口
	3. 定义类型别名
	*/
	var a MyInt
	a = 12
	fmt.Println(a.string())
}
