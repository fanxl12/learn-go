package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a = 12
	type IT int // 定义了int的别名
	var c IT = IT(a)
	fmt.Println(c)

	//字符串转数字
	var istr = "12"
	// go语言转换如果错误不会抛出异常，而是接收err错误信息
	myint, err := strconv.Atoi(istr)
	if err != nil {
		// 这里不是判断null，而是判断nil 如果不是nil，说明有err错误
		fmt.Println("convert to int failed")
	}
	fmt.Println(myint)

	// 字符串转换为flat32 转换成bool
	myf, err := strconv.ParseFloat("3.1514", 64)
	if err != nil {
		fmt.Println("convert to float failed")
	}
	fmt.Println(myf)

	parseBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println("convert to bool failed")
	}
	fmt.Println(parseBool)

	// 基本类型转字符串
	formatBool := strconv.FormatBool(true)
	fmt.Println(formatBool)
}
