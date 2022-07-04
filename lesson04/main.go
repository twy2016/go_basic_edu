package main

import "fmt"

const (
	BEIJING = iota * 5
	SHANGHAI
)

func main() {
	var a int
	fmt.Printf("a=%d,a的类型是:%T\n", a, a)

	var b int = 100
	fmt.Printf("b=%d,b的类型是:%T\n", b, b)

	var c = 100
	fmt.Printf("c=%d,c的类型是:%T\n", c, c)

	var d = "Go语言学习"
	fmt.Printf("d=%s,d的类型是:%T\n", d, d)

	// 短声明 :=
	e := 100
	fmt.Printf("e=%d,e的类型是:%T\n", e, e)

	f := "Go语言学习"
	fmt.Printf("f=%s,f的类型是:%T\n", f, f)

	// 多个变量声明
	var x, y = 100, 200
	var m, n = 100, "Hello"

	var (
		i = 100
		j = "go study"
	)
	fmt.Println(x, y, m, n, i, j)

	const length = 100
	//length = 200
	fmt.Println("length=", length)

	fmt.Println("BEIJING=", BEIJING)
	fmt.Println("SHANGHAI=", SHANGHAI)
}
