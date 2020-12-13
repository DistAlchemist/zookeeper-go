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
	"zookeepergo/replicalog"
)

var root *datatree.Datatree

func connecttoclient(clientset []network.Client, cP chan network.NetMessage) {
	var ResponseCli []*net.Conn
	cResCli := make(chan *net.Conn)
	go network.ConnectToClientrRes(clientset[0], cResCli)
	fmt.Println("test block")
	ResponseCli = append(ResponseCli, <-cResCli)
	fmt.Println("begin ResponseHandler")
	go network.ResponseHandler(ResponseCli[0], cP, 0)
}

//Listentoclient only listen to connect and message of client port,clientport is set to 8007
func Listentoclient(clientport int, cP chan network.NetMessage) {
	tcpaddr, err := net.ResolveTCPAddr("tcp", "localhost:"+strconv.Itoa(clientport))
	if err != nil {
		fmt.Println("tcp resolve error")
	}
	tcplisten, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		fmt.Println("tcp listen error")
	}
	for {
		fmt.Println("now listening to client port")
		tcpconn1, err := tcplisten.Accept()
		if err != nil {
			fmt.Println("tcp accept error")
		}
		fmt.Println("listen to client port build")
		readinfo := make([]byte, 30)
		fmt.Println("begin read from client")
		n, err := tcpconn1.Read(readinfo)
		if err != nil {
			fmt.Println("Client connection closed")
			continue
		}
		fmt.Println("receive:")
		cP <- network.MessageDealer(readinfo[:n])
		tcpconn1.Close()
	}
}

//Leader perform operations a leader is supposed to do
func Leader(Conn []*net.Conn) {
	cP := make(chan network.NetMessage)
	/*//connect to client
	fmt.Println("begin connect to client")
	clientset := Readclientcfg()
	connecttoclient(clientset, cP)*/
	fmt.Println("begin listen to client port")
	go Listentoclient(8007, cP)
	//load znode
	if root == nil {
		root = datatree.NewZnode()
	}

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
				str := datatree.LookZnode(Message.Str, root)
				network.SendOnetimeMessage(Message.Info, Message.Id, Message.Type, Message.Info, str)
				fmt.Printf("deal with message %d-%s\n", Message.Type, Message.Str)
			}
		case Message := <-network.CR:
			if Message.Type <= 3 {
				fmt.Printf("deal with message %d-%s\n", Message.Type, Message.Str)
				for i := 0; i < len(Conn); i++ {
					network.SendMessage(Conn[i], network.Winner, 3, network.Winner)
					time.Sleep(50 * time.Millisecond)
				}
			}
			if Message.Type == 9 {
				fmt.Printf("deal with message %d-%s\n", Message.Type, Message.Str)
				for j := 0; j < 2; j++ {
					if network.Peerset[j].Sid == Message.Id {
						for i := Message.Info + 1; i <= replicalog.Lognum; i++ {
							network.SendDataMessage(Conn[j], network.Winner, replicalog.Getlog(i).Action, replicalog.Getlog(i).Info, replicalog.Getlog(i).Str)
							time.Sleep(50 * time.Millisecond)
						}
					}
				}

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
