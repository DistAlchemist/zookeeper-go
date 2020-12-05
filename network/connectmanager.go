package network

import (
	"net"
)

//BeginConnect call all connection between servers
func BeginConnect(cConn chan *net.Conn, cRes chan *net.Conn, Peerset []Peer) {

	go ConnectToServer(Peerset[0], cConn)
	go ConnectToServer(Peerset[1], cConn)
	go ConnectToServerRes(Peerset[0], cRes)
	go ConnectToServerRes(Peerset[1], cRes)

}
