package quorum

import (
	"fmt"
	"zookeepergo/datatree"
	"zookeepergo/network"
	"zookeepergo/replicalog"
)

//Follower perform operations a follower is supposed to do
func Follower(cR chan network.NetMessage, Peerset []network.Peer, winner int, Sid int) {
	//load znode
	if root == nil {
		root = datatree.NewZnode()
	}
	fmt.Println("load a new node")
	//sync with leader
	for i := 0; i < 2; i++ {
		if Peerset[i].Sid == network.Winner {
			network.SendMessage(network.Conn[i], Sid, 9, replicalog.Lognow)
			fmt.Printf("sending sync information to %d\n", Peerset[i].Sid)
		}
	}

	//deal with message by select
	for {
		select {
		case Message := <-cR:
			fmt.Printf("peer: %d relative id,winner is %d\n", Message.Info, winner)
			if Message.Type == 4 {
				if Peerset[Message.Info].Sid == winner {
					fmt.Printf("peer: %d collapsed,winner is %d\n", Peerset[Message.Info].Sid, winner)
					network.Winner = -1
				}
				return
			}
			if Message.Type >= 5 {
				datatree.DealWithMessage(Message, root)
				fmt.Printf("deal with message %d-%s\n", Message.Type, Message.Str)
			}

		}
	}
}
