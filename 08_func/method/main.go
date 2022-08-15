package main

import "fmt"

// 方法，必须要有一个接收者
func main() {
	age := Age(25)
	age.String()
	age.Modify()
	age.String()
	age.Modify2()
	age.String()
}

type Age uint

func (age Age) String() {
	fmt.Println("the age is", age)
}

func (age *Age) Modify() {
	*age = Age(30)
}

func (age Age) Modify2() {
	age = Age(35)
}
