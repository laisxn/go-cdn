package main

import (
	"fmt"
	"go-cdn/config"
	"io"
	"net"
	"os"
	"time"
)

func main() {

	//创建一个tcp连接 ， 端口为8888

	listener, err := net.Listen("tcp", "127.0.0.1:"+config.Get("receive.port"))

	if err != nil {

		fmt.Println("net.Listen err ", err)

		return

	}

	//延迟关闭监听

	defer listener.Close()

	//堵塞等待用户连接， 可以用for 一直监听

	for {
		conn, err := listener.Accept()

		if err != nil {

			fmt.Println("listener.Accept err :", err)

			return

		}

		//延迟关闭
		defer conn.Close()

		go handleAccept(conn)

		time.Sleep(1)
	}

}

func handleAccept(conn net.Conn) {
	buf := make([]byte, 1024)

	//将文件内容缓存进入buf

	n, err := conn.Read(buf)

	if err != nil {

		fmt.Println("conn.Read err ", err)

		return

	}

	//将缓存区里面的文件名称赋值给fileName

	fileName := string(buf[:n])

	//回复“ok”

	conn.Write([]byte("ok"))

	//接收文件内容

	RecvFile(fileName, conn)
}

func RecvFile(fileName string, conn net.Conn) {

	//新建文件

	dir := "./cdn/"
	os.MkdirAll(dir, 755)

	f, err := os.Create(dir + fileName)

	if err != nil {

		fmt.Println("os.Create err:", err)

		return

	}

	//定义缓存字节切片

	buf := make([]byte, 1024*4)

	//接收多少 ， 写多少

	for {

		n, err := conn.Read(buf)

		if err != nil {

			if err == io.EOF {

				fmt.Println("文件接收完毕" + fileName)

			} else {

				fmt.Println("conn.Read err:", err)

			}

			return

		}

		if n == 0 {

			fmt.Println("n==0, 文件接收完毕")

			return

		}

		f.Write(buf[:n])

	}

}
