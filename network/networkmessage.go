package network

import "net"

type NetMessage struct {
	id int
	a  int //1:ask for elect,2:accept elect
}

var MeassageChannel chan NetMessage

func ListentoMessage(c net.Conn) {

	//c.Read()
}
