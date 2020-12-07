package network

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

//Relisten to the collapsed node and wait for reconnect
func Relisten(s Peer, id int) { ///TODO:change it to listen Dial will meet some problem on unopened socket
	fmt.Println("relisten to " + s.Addr + ":" + strconv.Itoa(s.Port))
	time.Sleep(1 * time.Second)
	Ctcplistener[id].Close()
	tcpaddr, err := net.ResolveTCPAddr("tcp", s.Addr+":"+strconv.Itoa(s.Port))
	Ctcplistener[id], err = net.ListenTCP("tcp", tcpaddr)
	tcpconn1, err := Ctcplistener[id].Accept()
	if err == nil {
		fmt.Println("connect to " + s.Addr + ":" + strconv.Itoa(s.Port))
		Conn[id] = &tcpconn1
	} else {
		return
	}
	time.Sleep(3 * time.Second)
	tcpconn, _ := net.DialTimeout("tcp", s.Addr+":"+strconv.Itoa(s.Portresponse), 2*time.Second)
	Response[id] = &tcpconn
	fmt.Printf("reconnect to %d successfully\n", id)
	go ResponseHandler(Response[id], CR, id)

}
