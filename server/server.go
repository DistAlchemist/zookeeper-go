package server

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"zookeepergo/network"
	"zookeepergo/quorum"
)

type Server struct {
	Sid         int
	CurrentVote int
}

func deleteNewLine(s string) string {
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, "\r", "", -1)
	return s
}

func deleteKuohao(s string) string {
	s = strings.Replace(s, "[", "", -1)
	s = strings.Replace(s, "]", "", -1)
	s = strings.Replace(s, ",", "", -1)
	return s
}

func (server *Server) readcfg() []network.Peer {
	var err error
	var byte []byte
	var peerset []network.Peer
	fp, err := os.Open("zoo.cfg")
	if err != nil {
		fmt.Println("open failed", err)
	}
	r := bufio.NewReader(fp)
	//read sid
	byte, err = r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
	}
	slice := string(byte)
	slice = deleteNewLine(slice)
	slice = strings.Split(slice, "=")[1]
	server.Sid, err = strconv.Atoi(slice)
	if err != nil {
		fmt.Println("string to int failed", err)
	}
	fmt.Printf("id of this station is %d \n", server.Sid)
	//read peer
	byte, err = r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
	}
	slice = string(byte)
	slice = deleteNewLine(slice)
	slice = strings.Split(slice, "=")[1]
	sidset := strings.Split(slice, ",")

	byte, err = r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
	}
	slice = string(byte)
	slice = deleteNewLine(slice)
	slice = strings.Split(slice, "=")[1]
	addrset := strings.Split(slice, ",")

	byte, err = r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
	}
	slice = string(byte)
	slice = deleteNewLine(slice)
	slice = strings.Split(slice, "=")[1]
	portset := strings.Split(slice, ",")

	byte, err = r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
	}
	slice = string(byte)
	slice = deleteNewLine(slice)
	slice = strings.Split(slice, "=")[1]
	portresponseset := strings.Split(slice, ",")

	var newpeer network.Peer
	for i := 0; i < len(sidset); i++ {
		peerset = append(peerset, newpeer)
		sidset[i] = deleteKuohao(sidset[i])
		peerset[i].Sid, err = strconv.Atoi(sidset[i])
		addrset[i] = deleteKuohao(addrset[i])
		peerset[i].Addr = addrset[i]
		portset[i] = deleteKuohao(portset[i])
		peerset[i].Port, err = strconv.Atoi(portset[i])

		portresponseset[i] = deleteKuohao(portresponseset[i])
		peerset[i].Portresponse, err = strconv.Atoi(portresponseset[i])
		fmt.Printf("read peer %d, %s:%d,%d\n", peerset[i].Sid, peerset[i].Addr, peerset[i].Port, peerset[i].Portresponse)
	}
	return peerset

}
func (server Server) Start() {
	fmt.Println("server begin")

	fmt.Println("load zoo.cfg")
	Peerset := server.readcfg()
	fmt.Printf("read a total of %d peer\n", len(Peerset))

	fmt.Println("initialize network")
	var Conn []*net.Conn
	var Response []*net.Conn
	cConn := make(chan *net.Conn)
	cRes := make(chan *net.Conn)
	network.BeginConnect(cConn, cRes, Peerset)
	for i := 0; i < 2; i++ {
		Conn = append(Conn, <-cConn)
		Response = append(Response, <-cRes)
	}
	fmt.Printf("%d %d\n", len(Conn), len(Response))

	fmt.Println("start election")
	quorum.LookForLeader(Peerset, server.Sid, Conn, Response)
	//winner := quorum.LookForLeader(Peerset, server.Sid, Conn, Response)
	/*if winner == server.Sid {
		quorum.Leader()
	} else {
		quorum.Follower()
	}*/
	fmt.Println("load znode")
}
