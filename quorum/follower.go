package quorum

import (
	"fmt"
	"zookeepergo/network"
)

//Follower perform operations a follower is supposed to do
func Follower(cR chan network.NetMessage, Peerset []network.Peer, winner int) {
	//load znode

	//sync with leader

	//deal with message by select
	for {
		select {
		case Message := <-cR:

			if Message.Type == 4 && Peerset[Message.Info].Sid == winner {
				fmt.Printf("peer: %d collapsed,winner is %d\n", Peerset[Message.Info].Sid, winner)
				return
			}
		}
	}
}
