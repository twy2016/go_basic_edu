package main

import "fmt"

func main() {
	//if
	i := 10
	if i > 10 {
		fmt.Println("i>10")
	} else {
		fmt.Println("i<=10")
	}

	if i := 6; i > 10 {
		fmt.Println("i>10")
	} else if i >= 6 && i <= 10 {
		fmt.Println("i>=6&&i<=10")
	} else {
		fmt.Println("i<10")
	}

	//switch
	switch i := 7; {
	case i > 10:
		fmt.Println("i>10")
	case i > 5 && i <= 10:
		fmt.Println("5<i<=10")
	default:
		fmt.Println("i<=5")
	}

	//默认每个case自带break,fallthrough
	switch j := 1; j {
	//case "a":
	case 1:
		fmt.Println(fmt.Sprintf("%s%d", "case1:", j))
		fallthrough
	case 2:
		fmt.Println(fmt.Sprintf("%s%d", "case2:", j))
	default:
		fmt.Println("没有匹配")
	}

	//for循环
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("the sum is", sum)

	//无限循环
	//for {
	//	fmt.Println("i")
	//}

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 2 {
				// 设置退出标签
				goto breakTag
			}
			fmt.Printf("%v-%v\n", i, j)
		}
	}
	return
breakTag:
	fmt.Println("结束for循环")
}
