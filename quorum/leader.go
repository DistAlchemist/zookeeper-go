package quorum

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"zookeepergo/datatree"
	"zookeepergo/network"
)

func connecttoclient(clientset []network.Client, cP chan network.NetMessage) {
	var ResponseCli []*net.Conn
	cResCli := make(chan *net.Conn)
	go network.ConnectToClientrRes(clientset[0], cResCli)
	fmt.Println("test block")
	ResponseCli = append(ResponseCli, <-cResCli)
	fmt.Println("begin ResponseHandler")
	go network.ResponseHandler(ResponseCli[0], cP, 0)
}

//Leader perform operations a leader is supposed to do
func Leader(Conn []*net.Conn) {
	//connect to client

	fmt.Println("begin connect to client")
	clientset := Readclientcfg()
	cP := make(chan network.NetMessage)
	connecttoclient(clientset, cP)

	//load znode
	root := datatree.NewZnode()
	fmt.Println("load a new node")
	//sync with follower

	//deal with message by select
	fmt.Println("begin dealing message")
	for {
		select {
		case Message := <-cP:
			if Message.Type >= 5 {
				datatree.DealWithMessage(Message, root)
				fmt.Printf("deal with message %d-%s\n", Message.Type, Message.Str)
				for i := 0; i < len(Conn); i++ {
					network.SendDataMessage(Conn[i], Message.Id, Message.Type, Message.Info, Message.Str)
					time.Sleep(50 * time.Millisecond)
				}
			}
			if Message.Type == 7 {
				datatree.LookZnode(Message.Str, root)
				fmt.Printf("deal with message %d-%s\n", Message.Type, Message.Str)
			}

		}
	}

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

//Readclientcfg read client config
func Readclientcfg() []network.Client {
	var err error
	var byte []byte
	var clientset []network.Client
	fp, err := os.Open("client.cfg")
	if err != nil {
		fmt.Println("open failed", err)
	}
	r := bufio.NewReader(fp)

	byte, err = r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
	}
	slice := string(byte)
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

	var newclient network.Client
	for i := 0; i < len(sidset); i++ {
		clientset = append(clientset, newclient)
		sidset[i] = deleteKuohao(sidset[i])
		clientset[i].Sid, err = strconv.Atoi(sidset[i])
		addrset[i] = deleteKuohao(addrset[i])
		clientset[i].Addr = addrset[i]
		portset[i] = deleteKuohao(portset[i])
		clientset[i].Port, err = strconv.Atoi(portset[i])
		portresponseset[i] = deleteKuohao(portresponseset[i])
		clientset[i].Portresponse, err = strconv.Atoi(portresponseset[i])
		fmt.Printf("read peer %d, %s:%d,%d\n", clientset[i].Sid, clientset[i].Addr, clientset[i].Port, clientset[i].Portresponse)
	}
	return clientset
}
