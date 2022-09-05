package send

import (
	"fmt"
	"go-cdn/config"
	"io"
	"net"
	"os"
)

func Send(path string) {

	file, err := os.Open(path)

	if err != nil {

		fmt.Println("os.Open err = ", err)

		return

	}

	//获取文件属性

	f, err := file.Stat()

	fmt.Println(f.Name(), f.Size())

	//主动连接服务器

	conn, err := net.Dial("tcp", config.Get("send.receive_ip")+":"+config.Get("send.receive_port"))

	if err != nil {

		fmt.Println("net.Dial err=", err)

		return

	}

	//文件接收完毕关闭连接

	defer conn.Close()

	//给接收方发送文件名

	_, err = conn.Write([]byte(f.Name()))

	if err != nil {

		fmt.Println("conn.Write err ", err)

		return

	}

	//接受对方回复， 如果回复“ok” 说明对方准备好了， 可以发送文件了

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)

	if err != nil {

		fmt.Println("conn.Read err=", err)

		return

	}

	//接收到ok发送文件内容

	if "ok" == string(buf[:n]) {

		//发送文件内容

		sendFile(path, conn)

	}
}

func sendFile(path string, conn net.Conn) {

	//以只读方式打开文件

	f, err := os.Open(path)

	if err != nil {

		fmt.Println("os.Open err:", err)

		return

	}

	//延迟关闭

	defer f.Close()

	//定义缓存字节切片

	buf := make([]byte, 1024*4)

	for {

		//从文件读取数据写入到buf缓存

		n, err := f.Read(buf)

		if err != nil {

			if err == io.EOF {

				fmt.Println("文件发送完毕")

			} else {

				fmt.Println("f.Read err:", err)

			}

			return

		}

		//发送内容

		conn.Write(buf[:n])

	}

}
