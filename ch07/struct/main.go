package main

import "fmt"

type Person struct {
	name   string
	age    int
	addr   string
	height int
	sex    string
}

func main() {
	// 如何初始化结构体
	p1 := Person{"body1", 18, "武汉", 180, "女"}
	fmt.Println(p1)
	p2 := Person{
		name:   "body2",
		age:    18,
		addr:   "上海",
		height: 170,
		sex:    "女",
	}
	fmt.Println(p2)

	//匿名结构体
	address := struct {
		province string
		city     string
		address  string
	}{
		"北京市",
		"通州区",
		"xxxx",
	}
	fmt.Println(address.city)
}
