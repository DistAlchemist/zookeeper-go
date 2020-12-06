package quorum

import (
	"fmt"
	"zookeepergo/datatree"
	"zookeepergo/network"
)

//Follower perform operations a follower is supposed to do
func Follower(cR chan network.NetMessage, Peerset []network.Peer, winner int) {
	//load znode
	if root == nil {
		root = datatree.NewZnode()
	}
	fmt.Println("load a new node")
	//sync with leader

	//deal with message by select
	for {
		select {
		case Message := <-cR:
			fmt.Printf("peer: %d relative id,winner is %d\n", Message.Info, winner)
			if Message.Type == 4 && Peerset[Message.Info].Sid == winner {
				fmt.Printf("peer: %d collapsed,winner is %d\n", Peerset[Message.Info].Sid, winner)
				return
			}
			if Message.Type >= 5 {
				datatree.DealWithMessage(Message, root)
				fmt.Printf("deal with message %d-%s\n", Message.Type, Message.Str)
			}

		}
	}
}
