package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
	"zookeepergo/network"
)

//Responseport response port of this client
var Responseport int

func createsendsocket() *net.Conn {
	tcpconn, err := net.DialTimeout("tcp", "localhost:8007", 2*time.Second)
	if err != nil {
		fmt.Println("client connect error ", err)
		time.Sleep(300 * time.Millisecond)
	}
	return &tcpconn
}

var tcplisten net.Listener

//Start of client
func Start() {
	var byte []byte
	fp, err := os.Open("clientport.cfg")
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
	Responseport, _ = strconv.Atoi(slice)
	fmt.Printf("clientport:%d\n", Responseport)
	tcplisten, _ = net.Listen("tcp", "localhost:"+strconv.Itoa(Responseport))
	for {

		var strinput1, strinput2 string
		fmt.Println("please input command: ")
		fmt.Scanf("%s\n", &strinput1)
		time.Sleep(50 * time.Millisecond)

		if strinput1 == "CREATE" {
			fmt.Println("please input create dir ")
			fmt.Scanf("%s\n", &strinput2)
			tcpconn := *createsendsocket()
			Create(strinput2, tcpconn)
			tcpconn.Close()
		} else if strinput1 == "DELETE" {
			fmt.Println("please input delete dir ")
			fmt.Scanf("%s\n", &strinput2)
			tcpconn := *createsendsocket()
			Delete(strinput2, tcpconn)
			tcpconn.Close()
		} else if strinput1 == "DIR" {
			fmt.Println("please input show dir ")
			fmt.Scanf("%s\n", &strinput2)
			tcpconn := *createsendsocket()
			Dir(strinput2, tcpconn)
			tcpconn.Close()
		} else if strinput1 == "WATCH" {
			var watcherport int
			fmt.Println("please input watch dir ")
			fmt.Scanf("%s\n", &strinput2)
			fmt.Println("please input watcher port ")
			fmt.Scanf("%d\n", &watcherport)
			tcpconn := *createsendsocket()
			Watch(strinput2, tcpconn, watcherport, nil)
			tcpconn.Close()
		} else {
			fmt.Println("input is not a command ")
		}
		time.Sleep(100 * time.Millisecond)

	}
}

//ListentoWatcherport listen to watcherport to get watchevent
func ListentoWatcherport(watcherport int, dest *string) {
	tcplisten1, _ := net.Listen("tcp", "localhost:"+strconv.Itoa(watcherport))
	fmt.Println("now listen to " + strconv.Itoa(watcherport))
	tcpconn1, err := tcplisten1.Accept()
	if err != nil {
		fmt.Println("receive watchedevent error")
		tcpconn1.Close()
		return
	}
	readinfo := make([]byte, 30)
	fmt.Println("begin read from client")
	n, _ := tcpconn1.Read(readinfo)
	fmt.Println("watchedevent:" + network.MessageDealer(readinfo[:n]).Str)
	tcpconn1.Close()
	if dest != nil {
		*dest = network.MessageDealer(readinfo[:n]).Str
	}
	return
}
