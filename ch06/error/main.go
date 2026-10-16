package main

import (
	"errors"
	"fmt"
)

func A() (int, error) {
	return 0, errors.New("this is an error")
}

func B() (int, error) {
	panic("this is an panic") //panic会导致程序的退出， 平时开发的时候使用很少
	return 0, errors.New("this is an error")
}

func C() (int, error) {
	defer func() {
		// recover 这个函数能捕获panic
		if r := recover(); r != nil {
			fmt.Println("recovered:", r)
		}
	}()
	var names map[string]string
	names["go"] = "go工程师" // 由于names没初始化 这里会抛panic错误
	return 0, errors.New("this is an error")

	// 1.defer需要放在panic之前定义，另外 recover 只有在defer调用的函数中才会生效
	// 2.recover 处理异常后，逻辑并不会恢复到panic的那个点去
	// 3.多个defer会形成栈，后定义的defer会先执行
}

func main() {
	if _, err := A(); err != nil {
		fmt.Println(err)
	}
}
