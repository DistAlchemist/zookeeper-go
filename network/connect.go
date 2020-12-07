package network

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

//Conn global connection
var Conn []*net.Conn

//Response global response
var Response []*net.Conn

//Ctcplistener tcp connect listener
var Ctcplistener []*net.TCPListener

//Rtcplistener tcp response listener
var Rtcplistener []*net.TCPListener

//ConnectToServer build a connection to a server
func ConnectToServer(num int, a chan *net.Conn, c chan int) {
	s := Peerset[num]
	tcpaddr, err := net.ResolveTCPAddr("tcp", s.Addr+":"+strconv.Itoa(s.Port))
	tcpconn, err := net.DialTimeout("tcp", s.Addr+":"+strconv.Itoa(s.Port), 2*time.Second)
	fmt.Println("send TCP request to " + s.Addr + ":" + strconv.Itoa(s.Port))
	if err != nil {
		fmt.Println("send TCP request failed, listen to " + s.Addr + ":" + strconv.Itoa(s.Port))
		Ctcplistener[num], err = net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			fmt.Println("listen error")
		}
		tcpconn1, err := Ctcplistener[num].Accept()
		if err != nil {
			fmt.Println("accept error")
		}
		fmt.Println("connect to " + s.Addr + ":" + strconv.Itoa(s.Port))
		/*readinfo := make([]byte, 10)
		tcpconn1.Read(readinfo)
		fmt.Println("reveived:" + string(readinfo) + " from " + s.Addr + ":" + strconv.Itoa(s.Port))*/
		a <- &tcpconn1
		c <- 1
	} else {
		fmt.Println("connect to " + s.Addr + ":" + strconv.Itoa(s.Port))
		/*fmt.Println("send test to " + s.Addr + ":" + strconv.Itoa(s.Port))
		tcpconn.Write([]byte("send test"))*/
		a <- &tcpconn
		c <- 1
	}
}

//ConnectToServerRes connect to server response port to get response
func ConnectToServerRes(num int, a chan *net.Conn, c chan int) {
	s := Peerset[num]
	tcpaddr, err := net.ResolveTCPAddr("tcp", s.Addr+":"+strconv.Itoa(s.Portresponse))
	tcpconn, err := net.DialTimeout("tcp", s.Addr+":"+strconv.Itoa(s.Portresponse), 2*time.Second)
	fmt.Println("send TCP request to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))
	if err != nil {
		fmt.Println("send TCP request failed, listen to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))
		Rtcplistener[num], err = net.ListenTCP("tcp", tcpaddr)
		if err != nil {
			fmt.Println("listen error")
		}
		tcpconn1, err := Rtcplistener[num].Accept()
		if err != nil {
			fmt.Println("accept error")
		}
		fmt.Println("connect to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))
		/*readinfo := make([]byte, 10)
		tcpconn1.Read(readinfo)
		fmt.Println("reveived:" + string(readinfo) + " from " + s.Addr + ":" + strconv.Itoa(s.Portresponse))*/
		a <- &tcpconn1
		c <- 1
	} else {
		fmt.Println("connect to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))
		/*fmt.Println("send response test to " + s.Addr + ":" + strconv.Itoa(s.Portresponse))
		tcpconn.Write([]byte("response test"))*/
		a <- &tcpconn
		c <- 1

	}
}
