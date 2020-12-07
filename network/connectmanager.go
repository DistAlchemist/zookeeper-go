package network

import (
	"net"
)

//BeginConnect call all connection between servers
func BeginConnect(cConn chan *net.Conn, cRes chan *net.Conn, Peerset []Peer) {
	Conn1 := make([]*net.Conn, len(Peerset))
	Res1 := make([]*net.Conn, len(Peerset))
	c := make(chan int)
	cC0 := make(chan *net.Conn)
	cC1 := make(chan *net.Conn)
	cR0 := make(chan *net.Conn)
	cR1 := make(chan *net.Conn)
	go ConnectToServer(0, cC0, c)
	go ConnectToServer(1, cC1, c)
	go ConnectToServerRes(0, cR0, c)
	go ConnectToServerRes(1, cR1, c)
	//wait until all nodes ready, this is a feature that assume all nodes must work well when starting
	Conn1[0] = <-cC0
	Conn1[1] = <-cC1
	Res1[0] = <-cR0
	Res1[1] = <-cR1
	for i := 0; i < 2*len(Peerset); i++ {
		<-c
	}
	/*_, err := (*Conn1[0]).Write([]byte("1234135"))
	if err != nil {
		fmt.Println(err)
	}*/
	for i := 0; i < len(Peerset); i++ {
		cConn <- Conn1[i]
		cRes <- Res1[i]
	}
	//this complex process is to make async connect process in order
}
