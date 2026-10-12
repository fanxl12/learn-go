package main

import "fmt"

// 全局变量 和 局部变量
// 全局变量是可以允许不使用的
var (
	name1 = "body"
	age   = 18
	ok    bool
)

func main() {
	//go是静态语言，三种定义变量的方式
	var name int
	name = 1
	fmt.Println(name)

	var age = 2
	fmt.Println(age)

	//简洁变量定义 全局不能使用简洁变量定义方式
	school := "School"
	fmt.Println(school)

	//go语言中 变量定义了不使用是不行的

	//2.多变量定义
	var user1, user2, user3 = "body1", 1, "body2"
	fmt.Println(user1, user2, user3)

	/**
	变量必须先定义才能使用
	go语言是静态语言
	变量名不能冲突
	变量有默认值
	*/
}
