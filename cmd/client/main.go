package main

import (
	"fmt"
	"net"
	"time"
	"zookeepergo/network"
)

func main() {
	tcpconn, err := net.DialTimeout("tcp", "localhost:8007", 2*time.Second)
	if err != nil {
		fmt.Println("client connect error ", err)
	}
	for {

		var strinput1, strinput2 string
		fmt.Println("please input command: ")
		fmt.Scanf("%s\n", &strinput1)
		time.Sleep(50 * time.Millisecond)
		if strinput1 == "CREATE" {
			fmt.Println("please input create dir ")
			fmt.Scanf("%s\n", &strinput2)
			network.SendDataMessage(&tcpconn, 4, 5, 1, strinput2)
		} else if strinput1 == "DELETE" {
			fmt.Println("please input delete dir ")
			fmt.Scanf("%s\n", &strinput2)
			network.SendDataMessage(&tcpconn, 4, 6, 1, strinput2)
		} else if strinput1 == "DIR" {
			fmt.Println("please input show dir ")
			fmt.Scanf("%s\n", &strinput2)
			network.SendDataMessage(&tcpconn, 4, 7, 1, strinput2)
		} else if strinput1 == "WATCH" {
			fmt.Println("please input watch dir ")
			fmt.Scanf("%s\n", &strinput2)
			network.SendDataMessage(&tcpconn, 4, 9, 1, strinput2)
		} else {
			fmt.Println("input is not a command ")
		}
		time.Sleep(100 * time.Millisecond)

	}
}
