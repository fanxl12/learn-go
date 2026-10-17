package main

import "fmt"

// 接口的定义
type Duck interface {
	//方法的申明
	Gaga()
	Walk()
	Swimming()
}

type pskDuck struct {
	legs int
}

func (p *pskDuck) Gaga() {
	fmt.Println("嘎嘎")
}

func (p *pskDuck) Walk() {
	fmt.Println("this is psk duck Walk")
}

func (p *pskDuck) Swimming() {
	fmt.Println("this is psk duck Swimming")
}

func main() {

	var d Duck = &pskDuck{}
	d.Gaga()
	d.Walk()
	d.Swimming()

}
