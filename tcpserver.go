package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn){
	defer conn.Close()
	for {
		var buf [128]byte
		n ,err := conn.Read(buf[:])
		conn.Write([]byte("receive"))
		if err != nil {
			fmt.Println("Read from tcp server failed,err:",err)
			break
		}
		data := string(buf[:n])
		fmt.Printf("Recived from client,data:%s\n",data)
	}
}

func main() {
	// 监听TCP 服务端口
	listener,err := net.Listen("tcp","0.0.0.0:9090")
	if err != nil {
		fmt.Println("Listen tcp server failed,err:",err)
		return
	}

	for{
		// 建立socket连接
		conn,err := listener.Accept()
		if err != nil {
			fmt.Println("Listen.Accept failed,err:",err)
			continue
		}

		// 业务处理逻辑
		go process(conn)
	}
}