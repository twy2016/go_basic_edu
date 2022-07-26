package main

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}

// 甩干器
type dryer struct {
}

// 海尔洗衣机
type haier struct {
	dryer
}

func (d dryer) dry() {
	fmt.Println("甩一甩")
}

func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	h := haier{
		dryer{},
	}
	h.wash()
	h.dry()
}
