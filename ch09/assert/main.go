package main

import (
	"fmt"
	"strings"
)

func add(a, b interface{}) interface{} {
	switch a.(type) {
	case int:
		return a.(int) + b.(int)
	case int32:
		ai, _ := a.(int32)
		bi, _ := b.(int32)
		return ai + bi
	case string:
		as, _ := a.(string)
		bs, _ := b.(string)
		return as + bs
	default:
		panic("not supported")
	}
}

func main() {
	a := 1
	b := 2
	fmt.Println(add(a, b))
	re := add("hello", " world")
	res, _ := re.(string)
	fmt.Println(strings.Split(res, " "))
	fmt.Println(res)
}
