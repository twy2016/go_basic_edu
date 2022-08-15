package main

import "fmt"

func main() {
	array := [5]string{"a", "b", "c", "d", "e"}
	//基于数组生成切片，包含索引start，但是不包含索引end
	//slice := array[start:end]
	slice := array[2:4]
	fmt.Println(slice)
	fmt.Printf("切片len:%v,切片cap:%v\n", len(slice), cap(slice))

	/*
		start 和 end 索引都是可以省略的
		array[:4] 等价于 array[0:4]
		array[1:] 等价于 array[1:5]
		array[:] 等价于 array[0:5]
	*/
	slice2 := array[:4]
	fmt.Println(slice2)
	slice3 := array[1:]
	fmt.Println(slice3)
	slice4 := array[:]
	fmt.Println(slice4)

	//切片修改
	slice[1] = "f"
	fmt.Println(slice)

	//切片声明
	slice5 := make([]string, 4, 8)
	fmt.Printf("切片：%v,切片len:%v,切片cap:%v\n", slice5, len(slice5), cap(slice5))

	//append方法 自动扩容
	slice6 := append(slice5, "m", "n", "x", "y", "z")
	fmt.Printf("切片：%v,切片len:%v,切片cap:%v\n", slice6, len(slice6), cap(slice6))

	//遍历
	for _, v := range slice6 {
		if v != "" {
			fmt.Printf("切片值:%v\n", v)
		}
	}
}
