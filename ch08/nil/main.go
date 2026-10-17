package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {

	/*
		不同的数据类型零值不一样
		bool      false
		numbers   0
		string    ""
		pointer   nil
		slice     nil
		map       nil
		channel/interface/function  nil

		struct  默认值不是nil，默认值是具体字段的默认值

	*/

	p1 := Person{
		name: "body",
		age:  20,
	}

	p2 := Person{
		name: "body",
		age:  20,
	}

	p3 := Person{
		name: "body",
		age:  18,
	}
	if p1 == p2 {
		fmt.Println("p1 == p2")
	}
	if p1 == p3 {
		fmt.Println("p1 == p3")
	}

	//slice的默认值
	var ps []Person             // nil slice
	var ps2 = make([]Person, 0) // empty slice
	if ps == nil {
		fmt.Println("ps nil slice")
	}
	if ps2 == nil {
		fmt.Println("ps2 nil slice")
	}

	//map的默认值
	var m map[string]string             //nil map
	var m2 = make(map[string]string, 0) //empty map
	if m == nil {
		fmt.Println("m is nil map")
	}
	if m2 == nil {
		fmt.Println("m2 is nil map")
	}

}
