package main

import "fmt"

type Person struct {
	name string
}

// *Person 传的是Person的指针
func changeName(name string, p *Person) {
	p.name = name
}

func main() {
	p := Person{"Bob"}
	changeName("BobOne", &p)
	fmt.Println(p.name)
}
