package network

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

//Client save client information
type Client struct {
	Sid          int
	Addr         string
	Port         int
	Portresponse int
}

func ConnectToClientrRes(s Client, c1 chan *net.Conn) {
	tcpaddr, err := net.ResolveTCPAddr("tcp", s.Addr+":"+strconv.Itoa(s.Portresponse))
	tcpconn, err := net.DialTimeout("tcp", s.Addr+":"+strconv.Itoa(s.Portresponse), 2*time.Second)
	fmt.Println("send TCP request to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))
	if err != nil {
		fmt.Println("send TCP request failed, listen to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))
		tcplisten, err := net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			fmt.Println("listen error")
		}
		tcpconn1, err := tcplisten.Accept()
		if err != nil {
			fmt.Println("accept error")
		}
		c1 <- &tcpconn1
		fmt.Println("connect to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))
	} else {
		c1 <- &tcpconn
		fmt.Println("1connect to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))

	}
}
