package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"zookeepergo/network"
)

const responseport = 8008

func main() {
	tcpaddr, _ := net.ResolveTCPAddr("tcp", "localhost:"+strconv.Itoa(responseport))
	tcplisten, _ := net.ListenTCP("tcp", tcpaddr)
	for {
		tcpconn, err := net.DialTimeout("tcp", "localhost:8007", 2*time.Second)
		if err != nil {
			fmt.Println("client connect error ", err)
			time.Sleep(3000 * time.Millisecond)
			continue
		}
		var strinput1, strinput2 string
		fmt.Println("please input command: ")
		fmt.Scanf("%s\n", &strinput1)
		time.Sleep(50 * time.Millisecond)

		if strinput1 == "CREATE" {
			fmt.Println("please input create dir ")
			fmt.Scanf("%s\n", &strinput2)
			network.SendDataMessage(&tcpconn, 4, 5, 1, strinput2)
			tcpconn.Close()
		} else if strinput1 == "DELETE" {
			fmt.Println("please input delete dir ")
			fmt.Scanf("%s\n", &strinput2)
			network.SendDataMessage(&tcpconn, 4, 6, 1, strinput2)
			tcpconn.Close()
		} else if strinput1 == "DIR" {
			fmt.Println("please input show dir ")
			fmt.Scanf("%s\n", &strinput2)
			network.SendDataMessage(&tcpconn, 4, 7, responseport, strinput2)
			tcpconn.Close()
			time.Sleep(100 * time.Millisecond)
			tcpconn, err := tcplisten.Accept()
			if err != nil {
				fmt.Println("receive dir error")
			} else {
				readinfo := make([]byte, 30)
				fmt.Println("begin read from client")
				n, _ := tcpconn.Read(readinfo)
				fmt.Println("dir now have:" + network.MessageDealer(readinfo[:n]).Str)
			}
			tcpconn.Close()
		} else if strinput1 == "WATCH" {
			fmt.Println("please input watch dir ")
			fmt.Scanf("%s\n", &strinput2)
			network.SendDataMessage(&tcpconn, 4, 9, 1, strinput2)
			tcpconn.Close()
		} else {
			fmt.Println("input is not a command ")
			tcpconn.Close()
		}
		time.Sleep(100 * time.Millisecond)

	}
}
