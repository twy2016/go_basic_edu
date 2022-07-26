package main

import "fmt"

type Mover interface {
	Move()
}

type Dog struct{}

func (d Dog) Move() {
	fmt.Println("狗会动")
}

// Cat 猫结构体类型
type Cat struct{}

// Move 使用指针接收者定义Move方法实现Mover接口
func (c *Cat) Move() {
	fmt.Println("猫会动")
}

func main() {
	var x Mover
	var d1 = Dog{}
	x = d1
	x.Move()
	var d2 = &Dog{}
	x = d2
	x.Move()

	//var c1 = Cat{}  // c1是Cat类型
	var c2 = &Cat{} // c2是*Cat类型
	//x = c1          // 不能将c1当成Mover类型
	//x.Move()
	x = c2 // 可以将c2当成Mover类型
	x.Move()
}
