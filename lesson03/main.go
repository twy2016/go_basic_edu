package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Go语言学习")
	gid := os.Getegid()
	fmt.Println(gid)
}
