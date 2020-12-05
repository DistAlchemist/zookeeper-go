package server

import (
	"fmt"
	"net"
	"strings"
	"zookeepergo/network"
	"zookeepergo/quorum"
)

//Server used to be the server state, now only Sid works
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

//Start is the entrance for server
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
	fmt.Println("all node is ready")
	cR := make(chan network.NetMessage)
	go network.ResponseHandler(Response[0], cR, 0)
	go network.ResponseHandler(Response[1], cR, 1)
	//cR is received message queue and Conn is sending socket
	fmt.Println("load znode")
	fmt.Println("start election")
	//quorum.LookForLeader(Peerset, server.Sid, Conn, Response)
	var winner int
	for {
		winner = quorum.LookForLeader(Peerset, server.Sid, Conn, Response, cR)
		if winner == server.Sid {
			quorum.Leader(Conn)
		} else {
			quorum.Follower(cR, Peerset, winner)
		}
	}

}
