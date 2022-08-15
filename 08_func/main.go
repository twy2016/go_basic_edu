package main

import (
	"errors"
	"fmt"
)

var x = 0

func main() {
	result := sum(1, 2)
	fmt.Println(result)

	result2, err := sum2(-1, 2)
	fmt.Println(err)
	fmt.Println(result2)

	result3, err2 := sum3(-1, 2)
	fmt.Println(err2)
	fmt.Println(result3)

	s, err3 := sum4(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
	fmt.Println(err3)
	fmt.Println(s)

	//匿名函数
	anySum := func(a, b int) int {
		return a + b
	}
	fmt.Println(anySum(1, 2))

	//闭包
	f := closure()
	g := closure()

	// 同一个函数分别多次调用
	f() // i: 1
	g() // i: 1

	// 对返回的闭包对象多次调用
	g() // i: 2
	g() // i: 3

	//全局变量有影响
	m := closure2()
	n := closure2()
	m() // i:1
	n() // i:2

}

func sum(a int, b int) int {
	return a + b
}

//多返回值
func sum2(a int, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	return a + b, nil
}

//返回值命名
func sum3(a int, b int) (sum int, err error) {
	if a < 0 || b < 0 {
		return 0, errors.New("a或者b不能是负数")
	}
	sum = a + b
	err = nil
	return
}

//可变参数
func sum4(params ...int) (sum int, err error) {
	for _, v := range params {
		sum += v
	}
	return
}

func closure() func() int {
	i := 0
	return func() int {
		i++
		// 局部变量，同一个方法的i地址相同
		fmt.Println("i的地址:", &i)
		return i
	}
}

func closure2() func() int {
	return func() int {
		x++
		// 全局变量，不同的方法的x地址也相同
		fmt.Println("x的地址:", &x)
		return x
	}
}
