package network

import (
	"fmt"
	"net"
	"strconv"
)

//ConnectToServer build a connection to a server
func ConnectToServer(s Servername) {
	tcpaddr, err := net.ResolveTCPAddr("tcp", s.addr+":"+strconv.Itoa(s.port))
	tcpconn, err := net.DialTCP("tcp", nil, tcpaddr)
	fmt.Println("send TCP request to " + s.addr + ":" + strconv.Itoa(s.port))
	if err != nil {
		fmt.Println("send TCP request failed, listen to " + s.addr + ":" + strconv.Itoa(s.port))
		tcplisten, err := net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			fmt.Println("listen error")
		}
		tcpconn1, err := tcplisten.Accept()
		if err != nil {
			fmt.Println("accept error")
		}
		go ListentoMessage(tcpconn1)
	} else {
		go ListentoMessage(tcpconn)
	}
	fmt.Println("connect to " + s.addr + ":" + strconv.Itoa(s.port))

}
