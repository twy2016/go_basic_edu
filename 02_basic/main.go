package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

const name = "唐万言"

func main() {
	var i = 10
	fmt.Println(i)
	var f32 float32 = 2.1
	var f64 float64 = 2.14159
	fmt.Println(f32)
	fmt.Println(f64)
	fmt.Println(name)
	fmt.Printf("类型：%T\n", f32)
	fmt.Printf("%v\n", f32)
	fmt.Printf("%+v\n", f32)
	fmt.Printf("%#v\n", f32)

	//Itoa函数可以把一个 int 类型转为 string
	i2s := strconv.Itoa(i)
	//Atoi函数可以把一个 string 类型转为 int
	s2i, err := strconv.Atoi(i2s)
	fmt.Println(i2s, s2i, err)

	s1 := "Hello"
	//判断s1的前缀是否是H
	fmt.Println(strings.HasPrefix(s1, "H"))
	//在s1中查找字符串o
	fmt.Println(strings.Index(s1, "o"))
	//把s1全部转为大写
	fmt.Println(strings.ToUpper(s1))
	//把s1全部转为小写
	fmt.Println(strings.ToLower(s1))

	s := "Hello唐万言"
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Println(len(s))
	//UTF8 编码下，一个汉字对应三个字节，所以字符串 s 的长度其实是 14
	fmt.Println(s[0], s[1], s[13])
	fmt.Println(utf8.RuneCountInString(s))
	//遍历字符串，i 是索引，r 是 unicode 字符对应的 unicode 码点
	for i, r := range s {
		fmt.Println(i, r)
	}
}
