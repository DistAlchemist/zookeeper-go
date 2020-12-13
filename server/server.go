package server

import (
	"fmt"
	"net"
	"strings"
	"zookeepergo/network"
	"zookeepergo/quorum"
	"zookeepergo/replicalog"
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
	network.Peerset = server.readcfg()
	fmt.Printf("read a total of %d peer\n", len(network.Peerset))

	fmt.Println("initialize network")

	network.Conn = make([]*net.Conn, len(network.Peerset))
	network.Response = make([]*net.Conn, len(network.Peerset))
	network.Ctcplistener = make([]*net.TCPListener, len(network.Peerset))
	network.Rtcplistener = make([]*net.TCPListener, len(network.Peerset))
	cConn := make(chan *net.Conn)
	cRes := make(chan *net.Conn)
	go network.BeginConnect(cConn, cRes, network.Peerset)
	for i := 0; i < len(network.Peerset); i++ {
		network.Conn[i] = <-cConn
		network.Response[i] = <-cRes
	}
	/*_, err := (*Conn[0]).Write([]byte("1234135"))
	if err != nil {
		fmt.Println(err)
	}*/
	fmt.Println("all node is ready")
	network.CR = make(chan network.NetMessage)
	go network.ResponseHandler(network.Response[0], network.CR, 0)
	go network.ResponseHandler(network.Response[1], network.CR, 1)
	//cR is received message queue and Conn is sending socket
	fmt.Println("load znode")
	replicalog.Initlog()
	fmt.Println("start election")
	//quorum.LookForLeader(Peerset, server.Sid, Conn, Response)
	network.Winner = -1
	for {
		if network.Winner == -1 {
			network.Winner = quorum.LookForLeader(network.Peerset, server.Sid, network.Conn, network.Response, network.CR)
		}
		if network.Winner == server.Sid {
			quorum.Leader(network.Conn)
		} else {
			quorum.Follower(network.CR, network.Peerset, network.Winner, server.Sid)
		}
	}

}
