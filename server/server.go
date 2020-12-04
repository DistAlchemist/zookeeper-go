package server

import (
	"bufio"
	"fmt"
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

func (server Server) readcfg() []network.Peer {
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
	slice = strings.Replace(slice, "\n", "", -1)
	slice = strings.Replace(slice, "\r", "", -1)
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
	slice = strings.Replace(slice, "\n", "", -1)
	slice = strings.Replace(slice, "\r", "", -1)
	slice = strings.Split(slice, "=")[1]
	sidset := strings.Split(slice, ",")

	byte, err = r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
	}
	slice = string(byte)
	slice = strings.Replace(slice, "\n", "", -1)
	slice = strings.Replace(slice, "\r", "", -1)
	slice = strings.Split(slice, "=")[1]
	addrset := strings.Split(slice, ",")

	byte, err = r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
	}
	slice = string(byte)
	slice = strings.Replace(slice, "\n", "", -1)
	slice = strings.Replace(slice, "\r", "", -1)
	slice = strings.Split(slice, "=")[1]
	portset := strings.Split(slice, ",")
	var newpeer network.Peer
	for i := 0; i < len(sidset); i++ {
		peerset = append(peerset, newpeer)
		sidset[i] = strings.Replace(sidset[i], "[", "", -1)
		sidset[i] = strings.Replace(sidset[i], "]", "", -1)
		sidset[i] = strings.Replace(sidset[i], ",", "", -1)
		peerset[i].Sid, err = strconv.Atoi(sidset[i])
		addrset[i] = strings.Replace(addrset[i], "[", "", -1)
		addrset[i] = strings.Replace(addrset[i], "]", "", -1)
		addrset[i] = strings.Replace(addrset[i], ",", "", -1)
		peerset[i].Addr = addrset[i]
		portset[i] = strings.Replace(portset[i], "[", "", -1)
		portset[i] = strings.Replace(portset[i], "]", "", -1)
		portset[i] = strings.Replace(portset[i], ",", "", -1)
		peerset[i].Port, err = strconv.Atoi(portset[i])
		fmt.Printf("read peer %d, %s:%d\n", peerset[i].Sid, peerset[i].Addr, peerset[i].Port)
	}
	return peerset

}
func (server Server) Start() {
	fmt.Println("server begin")

	fmt.Println("load zoo.cfg")
	Peerset := server.readcfg()
	fmt.Printf("read a total of %d peer\n", len(Peerset))

	fmt.Println("initialize network")

	fmt.Println("load znode")
	fmt.Println("start election")
	quorum.LookForLeader(Peerset, server.Sid)
}
