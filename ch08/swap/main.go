package main

import "fmt"

func swapA(a, b *int) {
	a, b = b, a
}

func swapB(a, b *int) {
	//保存临时值
	t := *a

	// 将a地址的值改成b指向的值
	*a = *b

	//将b地址的值改成a指向的值
	*b = t
}

func main() {
	a, b := 3, 4
	swapA(&a, &b)
	fmt.Println(a, b)
	swapB(&a, &b)
	fmt.Println(a, b)
}
