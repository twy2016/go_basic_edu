package main

import "fmt"

func test01() {
	var arr [5]int
	arr[0] = 0
	arr[1] = 1
	arr[2] = 2
	arr[3] = 3
	arr[4] = 4
	fmt.Println(arr)
}

func test02() {
	var arr1 = [5]int{5, 10, 15, 20, 25}
	fmt.Println(arr1)

	arr2 := [5]int{5, 10, 15, 20, 25}
	fmt.Println(arr2)

	arr3 := [5]int{5, 10}
	fmt.Println(arr3)

	arr4 := [5]int{1: 5, 4: 100}
	fmt.Println(arr4)

	arr5 := [...]int{5, 10, 15, 20, 25, 30}
	fmt.Println(arr5)
}

func test03() {
	// 特别注意数组的长度是类型的一部分，所以 [3]int 和 [5]int 是不同的类型
	arr1 := [3]int{15, 20, 25}
	arr2 := [5]int{15, 20, 25, 30, 35}
	fmt.Printf("type of arr1 is %T\n", arr1)
	fmt.Printf("type of arr2 is %T\n", arr2)
}

func test04() {
	// 定义多维数组
	arr := [3][2]string{
		{"1", "Go语言-1"},
		{"2", "Go语言-2"},
		{"3", "Go语言-3"}}
	fmt.Println(arr)
}

func arrLength() {
	arr := [...]string{"Golang", "Java", "Python"}
	fmt.Println("数组的长度是:", len(arr))
}

func showArr() {
	arr := [...]string{"Golang", "Java", "Python"}
	for index, value := range arr {
		fmt.Printf("arr[%d]=%s\n", index, value)
	}

	for _, value := range arr {
		fmt.Printf("value=%s\n", value)
	}
}

func arrByValue() {
	arr := [...]string{"Golang", "Java", "Python"}
	arr2 := arr
	arr2[1] = "Golang"
	fmt.Println(arr)
	fmt.Println(arr2)
}

func main() {
	//test01()
	//test02()
	//test03()
	//test04()
	//arrLength()
	//showArr()
	arrByValue()
}
