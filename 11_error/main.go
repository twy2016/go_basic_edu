package main

import (
	"errors"
	"fmt"
	"os"
)

func OpenFile(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_RDONLY, 0)
}

type Info struct {
	id int64
}

func queryById(id int64) (*Info, error) {
	if id <= 0 {
		return nil, errors.New("无效的id")
	}
	return new(Info), nil
}

// OpError 自定义结构体类型
type OpError struct {
	Op string
}

// Error OpError 类型实现error接口
func (e *OpError) Error() string {
	return fmt.Sprintf("无权执行%s操作", e.Op)
}

func main() {
	//_, err := OpenFile("./xx.go")
	//if err != nil {
	//	fmt.Println("打开文件失败,error:", err)
	//	return
	//}

	//id, err := queryById(0)
	//fmt.Println(id, err)

	var opError OpError
	opError.Op = "hello error!"
	result := opError.Error()
	fmt.Println(result)
}
