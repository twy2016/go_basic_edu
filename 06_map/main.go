package main

import "fmt"

func main() {
	nameAgeMap := make(map[string]int)
	nameAgeMap["twy"] = 28
	nameAgeMap["wxx"] = 26
	fmt.Printf("nameAgeMap: %v\n", nameAgeMap)

	nameAgeMap2 := map[string]int{"gp": 18}
	fmt.Printf("nameAgeMap2: %v\n", nameAgeMap2)

	//获取
	age, ok := nameAgeMap["twy"]
	if ok {
		fmt.Printf("age: %v\n", age)
	}
	//删除
	delete(nameAgeMap, "twy")
	fmt.Printf("nameAgeMap: %v\n", nameAgeMap)

	//遍历
	for k, v := range nameAgeMap {
		fmt.Println("Key is", k, ",Value is", v)
	}
	fmt.Println(len(nameAgeMap))
}
