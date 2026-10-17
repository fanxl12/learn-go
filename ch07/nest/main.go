package main

import "fmt"

type Person struct {
	name string
	age  int
	addr struct {
		city string
	}
}

// 结构体绑定方法 值传递 会拷贝一个
func (p Person) print() {
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

// 有可能该函数中想要改变p的值，或者person对象很大 指针传递
func (p *Person) print2() {
	p.age = 19
	fmt.Printf("name: %s, age: %d\n", p.name, p.age)
}

func main() {
	p := Person{
		name: "Bob",
		age:  23,
	}
	p.print()
}
