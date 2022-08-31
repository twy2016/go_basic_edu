package main

import "fmt"

func main() {
	//指针取值
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	name := "唐万言"
	nameP := &name //取地址
	fmt.Println("name变量的值为:", name)
	fmt.Println("name变量的内存地址为:", nameP)

	//内置new方法,可以传递类型给它。它会返回对应的指针类型
	intP := new(int)
	fmt.Printf("类型：%T\n", intP)
	fmt.Printf("值：%v\n", intP)
	stringP := new(string)
	fmt.Printf("类型：%T\n", stringP)
	fmt.Printf("值：%v\n", stringP)

	//取指针的值
	nameV := *nameP
	fmt.Println("name变量的值为:", nameV)

	//修改指针
	*nameP = "wxx"
	fmt.Println("nameP指针指向的值为:", *nameP)
	fmt.Println("name变量的值为:", name)

	age := 18
	modifyAge(age)
	fmt.Println("age的值为:", age)
	modifyAge2(&age)
	fmt.Println("age的值为:", age)
}

func modifyAge(age int) {
	age = 20
}

func modifyAge2(age *int) {
	*age = 20
}
