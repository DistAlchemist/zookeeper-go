package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
	"zookeepergo/network"
)

const responseport = 8008

func createsendsocket() *net.Conn {
	tcpconn, err := net.DialTimeout("tcp", "localhost:8007", 2*time.Second)
	if err != nil {
		fmt.Println("client connect error ", err)
		time.Sleep(300 * time.Millisecond)
	}
	return &tcpconn
}
func main() {
	tcplisten, _ := net.Listen("tcp", "localhost:"+strconv.Itoa(responseport))
	for {

		var strinput1, strinput2 string
		fmt.Println("please input command: ")
		fmt.Scanf("%s\n", &strinput1)
		time.Sleep(50 * time.Millisecond)

		if strinput1 == "CREATE" {
			fmt.Println("please input create dir ")
			fmt.Scanf("%s\n", &strinput2)
			tcpconn := *createsendsocket()
			network.SendDataMessage(&tcpconn, 4, 5, 1, strinput2)
			tcpconn.Close()
		} else if strinput1 == "DELETE" {
			fmt.Println("please input delete dir ")
			fmt.Scanf("%s\n", &strinput2)
			tcpconn := *createsendsocket()
			network.SendDataMessage(&tcpconn, 4, 6, 1, strinput2)
			tcpconn.Close()
		} else if strinput1 == "DIR" {
			fmt.Println("please input show dir ")
			fmt.Scanf("%s\n", &strinput2)
			tcpconn := *createsendsocket()
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
			var watcherport int
			fmt.Println("please input watch dir ")
			fmt.Scanf("%s\n", &strinput2)
			fmt.Println("please input watcher port ")
			fmt.Scanf("%d\n", &watcherport)
			tcpconn := *createsendsocket()
			network.SendDataMessage(&tcpconn, 4, 8, watcherport, strinput2)
			go ListentoWatcherport(watcherport)
			tcpconn.Close()
		} else {
			fmt.Println("input is not a command ")
		}
		time.Sleep(100 * time.Millisecond)

	}
}

//ListentoWatcherport listen to watcherport to get watchevent
func ListentoWatcherport(watcherport int) {
	tcplisten1, _ := net.Listen("tcp", "localhost:"+strconv.Itoa(watcherport))
	fmt.Println("now listen to " + strconv.Itoa(watcherport))
	tcpconn1, err := tcplisten1.Accept()
	if err != nil {
		fmt.Println("receive watchedevent error")
	} else {
		readinfo := make([]byte, 30)
		fmt.Println("begin read from client")
		n, _ := tcpconn1.Read(readinfo)
		fmt.Println("watchedevent:" + network.MessageDealer(readinfo[:n]).Str)
	}

	tcpconn1.Close()
}
