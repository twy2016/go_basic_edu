package main

/**
使用接口进行扩展
*/
import "fmt"

type Sayer interface {
	Say()
}

type Cat struct{}

type Dog struct{}

func (cat Cat) Say() {
	fmt.Println("喵喵喵")
}

func (dog Dog) Say() {
	fmt.Println("汪汪汪")
}

// MakeHungry 构造函数
func MakeHungry(s Sayer) {
	s.Say()
}

func main() {
	cat := Cat{}
	dog := Dog{}
	MakeHungry(cat)
	MakeHungry(dog)

	// 接口变量
	var x Sayer // 声明一个Sayer类型的变量x
	a := Cat{}  // 声明一个Cat类型变量a
	b := Dog{}  // 声明一个Dog类型变量b
	x = a       // 可以把Cat类型变量直接赋值给x
	x.Say()     // 喵喵喵
	x = b       // 可以把Dog类型变量直接赋值给x
	x.Say()     // 汪汪汪
}
