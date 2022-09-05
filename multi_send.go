package main

import (
	"fmt"
	"go-cdn/send"
	"io/ioutil"
	"log"
	"path/filepath"
)

func main() {

	//获取文件属性

	fmt.Println("请输入文件目录：")

	var dir string

	fmt.Scan(&dir)

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		send.Send(filepath.Join(dir, file.Name()))
	}

}
