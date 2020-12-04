package network

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

//ConnectToServer build a connection to a server
func ConnectToServer(s Peer) net.Conn {
	tcpaddr, err := net.ResolveTCPAddr("tcp", s.Addr+":"+strconv.Itoa(s.Port))
	tcpconn, err := net.DialTimeout("tcp", s.Addr+":"+strconv.Itoa(s.Port), 2*time.Second)
	fmt.Println("send TCP request to " + s.Addr + ":" + strconv.Itoa(s.Port))
	if err != nil {
		fmt.Println("send TCP request failed, listen to " + s.Addr + ":" + strconv.Itoa(s.Port))
		tcplisten, err := net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			fmt.Println("listen error")
		}
		tcpconn1, err := tcplisten.Accept()
		if err != nil {
			fmt.Println("accept error")
		}
		fmt.Println("connect to " + s.Addr + ":" + strconv.Itoa(s.Port))
		return tcpconn1
	} else {
		fmt.Println("connect to " + s.Addr + ":" + strconv.Itoa(s.Port))
		return tcpconn
	}
	return tcpconn
}
