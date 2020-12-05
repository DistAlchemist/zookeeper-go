package network

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

//NetMessage in the form of "Id:Type:Info"
type NetMessage struct {
	//note: id is relative
	Id   int
	Type int
	//Type=1 : ask for vote info=xid
	//Type=2 ：vote for  info=nowVote
	//Type=3 : i become winner, other ack
	Info int
}

//MessageFolder transalte message struct to byte
func MessageFolder(id int, typ int, info int) []byte {
	string1 := strconv.Itoa(id) + ":" + strconv.Itoa(typ) + ":" + strconv.Itoa(info)
	return []byte(string1)
}

//SendMessage send message
func SendMessage(c *net.Conn, Sid int, typ int, info int) {
	fmt.Printf("send %d:%d:%d\n", Sid, typ, info)
	(*c).Write(MessageFolder(Sid, typ, info))
}

//MessageDealer transalte message bytes to struct
func MessageDealer(bytes []byte) NetMessage {
	string1 := string(bytes)
	fmt.Println("receive: " + string1)
	stringsplit := strings.Split(string1, ":")
	var newNetMessage NetMessage
	var err error
	newNetMessage.Id, err = strconv.Atoi(strings.Replace(stringsplit[0], ":", "", -1))
	newNetMessage.Type, err = strconv.Atoi(strings.Replace(stringsplit[1], ":", "", -1))
	newNetMessage.Info, err = strconv.Atoi(strings.Replace(stringsplit[2], ":", "", -1))
	if err != nil {
		fmt.Println("MessageDealer Error")
	}
	return newNetMessage
}

//ResponseHandler handle information from other node when counting down
func ResponseHandler(c *net.Conn, ch chan NetMessage, num int) {
	for {
		readinfo := make([]byte, 30)
		n, err := (*c).Read(readinfo)
		if err != nil {
			fmt.Println("ResponseHandler Error " + string(readinfo[:n]))
			fmt.Printf("peer: %d collapsed\n", num)
			(*c).Close()
		} else {
			MessageProcessed := MessageDealer(readinfo[:n])
			MessageProcessed.Id = num
			ch <- MessageProcessed
		}
	}

	//c.Read()
}
