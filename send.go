package main

import (
	"fmt"
	"go-cdn/send"
)

func main() {

	//获取文件属性

	fmt.Println("请输入文件：")

	var path string

	fmt.Scan(&path)

	send.Send(path)
}
