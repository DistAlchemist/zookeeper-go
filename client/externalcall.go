package client

import (
	"fmt"
	"net"
	"time"
	"zookeepergo/network"
)

//this file is some external call function for go client

//Create create function of client
func Create(strinput2 string, tc net.Conn) {
	network.SendDataMessage(&tc, 4, 5, 1, strinput2)
}

//Delete delete function of client
func Delete(strinput2 string, tc net.Conn) {
	network.SendDataMessage(&tc, 4, 6, 1, strinput2)
}

//Dir return dir string
func Dir(strinput2 string, tc net.Conn) string {
	network.SendDataMessage(&tc, 4, 7, responseport, strinput2)
	time.Sleep(100 * time.Millisecond)
	tcpconn1, err := tcplisten.Accept()
	if err != nil {
		fmt.Println("receive dir error")
		return ""
	}
	readinfo := make([]byte, 30)
	fmt.Println("begin read from client")
	n, _ := tcpconn1.Read(readinfo)
	tcpconn1.Close()
	fmt.Println("dir now have:" + network.MessageDealer(readinfo[:n]).Str)
	return network.MessageDealer(readinfo[:n]).Str
}

func Watch(strinput2 string, tc net.Conn, watcherport int, dest *string) {
	network.SendDataMessage(&tc, 4, 8, watcherport, strinput2)
	go ListentoWatcherport(watcherport, dest)
}
