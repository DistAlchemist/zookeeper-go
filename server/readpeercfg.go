package server

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"zookeepergo/network"
)

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
