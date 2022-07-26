package main

import "fmt"

// Payer 包含支付方法的接口类型
type Payer interface {
	Pay(int64)
}

type Zhifubao struct {
	//支付宝
}

// Pay 支付宝的支付方法
func (z *Zhifubao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款：%.2f元。\n", float64(amount/100))
}

type WeChat struct {
	// 微信
}

// Pay 微信的支付方法
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款：%.2f元。\n", float64(amount/100))
}

//func Checkout(obj *Zhifubao) {
//	obj.Pay(100)
//}

// Checkout 结账
func Checkout(obj Payer) {
	// 支付100元
	obj.Pay(100)
}

func main() {
	Checkout(&Zhifubao{})
	Checkout(&WeChat{}) // 现在支持使用微信支付
}
