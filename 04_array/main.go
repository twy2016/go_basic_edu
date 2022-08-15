package main

import "fmt"

func main() {
	var arr = [5]string{"1", "2", "3", "4", "5"}
	fmt.Println(arr)
	fmt.Printf("%v\n", arr)
	fmt.Printf("%+v\n", arr)
	fmt.Printf("%#v\n", arr)

	array := [...]string{"a", "b", "c", "d", "e"}
	fmt.Println(len(array))

	for k, v := range array {
		fmt.Printf("对应索引:%d\n", k)
		fmt.Printf("对应值:%s\n", v)

	}
}
