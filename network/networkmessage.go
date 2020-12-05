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
	//Type=4 : node id collapsed, need to elect
	//Type=5 : CREATE
	//Type=6 : DELETE
	//Type=7 : DIR
	Info int
	Str  string
}

//MessageFolder transalte message struct to byte
func MessageFolder(id int, typ int, info int) []byte {
	string1 := strconv.Itoa(id) + ":" + strconv.Itoa(typ) + ":" + strconv.Itoa(info)
	return []byte(string1)
}

//MessageDataFolder transalte message struct to byte
func MessageDataFolder(id int, typ int, info int, str string) []byte {
	string1 := strconv.Itoa(id) + ":" + strconv.Itoa(typ) + ":" + strconv.Itoa(info) + ":" + str
	return []byte(string1)
}

//SendMessage send message
func SendMessage(c *net.Conn, Sid int, typ int, info int) {
	fmt.Printf("send %d:%d:%d\n", Sid, typ, info)
	(*c).Write(MessageFolder(Sid, typ, info))
}

//SendDataMessage send message
func SendDataMessage(c *net.Conn, Sid int, typ int, info int, str string) {
	fmt.Printf("send %d:%d:%d %s\n", Sid, typ, info, str)
	(*c).Write(MessageDataFolder(Sid, typ, info, str))
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
	if newNetMessage.Type >= 5 {
		newNetMessage.Str = strings.Replace(stringsplit[3], ":", "", -1)
	}
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
			(*c).Close()
			// when type=4 id is meaningless, only to tell that a node collapsed
			var MessageProcessed NetMessage
			MessageProcessed.Id = num
			MessageProcessed.Type = 4
			MessageProcessed.Info = num
			ch <- MessageProcessed
			return
		}
		MessageProcessed := MessageDealer(readinfo[:n])
		MessageProcessed.Id = num
		ch <- MessageProcessed
	}

	//c.Read()
}
