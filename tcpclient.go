package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	// 连接服务器
	conn,err := net.Dial("tcp","127.0.0.1:9090")
	if err != nil {
		fmt.Println("Connect to TCP server failed ,err:",err)
		return
	}

	// 读取命令行输入
	inputReader := bufio.NewReader(os.Stdin)

	// 一直读取直到遇到换行符
	for {
		input,err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Read from console failed,err:",err)
			return
		}

		// 读取到字符"Q"退出
		str := strings.TrimSpace(input)
		if str == "Q"{
			break
		}

		input = "wwwwwwwwwwwwwwwwwwwwwwwww" +
			"wwwwwwwwwwwwwwwwwwwwwwwwwwwww" +
			"wwwwwwwwwwwwwwwwfwfwwwwwwwwwww" +
			"wwwwwwwwwwwwwwwwwwwwwwwwwwww" +
			"wwwwwwwwwwwwwwwwwwwwwwwwwwww" +
			"wwwwwwwwwwwwwwwwwwwwwwwwwww" +
			"wwwwwwwwwwwwwwwwwwwwwwww"

		outStr := make([]byte,1)

		// 响应服务端信息
		 start := time.Now();

		_,err = conn.Write([]byte(input))
		conn.Read(outStr)

		cost  := time.Since(start)
		fmt.Printf("time cost = %v\n", cost)

		if err != nil{
			fmt.Println("Write failed,err:",err)
			break
		}
	}

}