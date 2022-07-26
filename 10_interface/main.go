package main

import "fmt"

// Singer 接口
type Singer interface {
	Sing()
}

type Bird struct{}

// Sing Bird类型的Sing方法
func (b Bird) Sing() {
	fmt.Println("唧唧唧")
}

func main() {
	bird := Bird{}
	bird.Sing()
}
