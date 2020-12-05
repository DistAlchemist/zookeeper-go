package quorum

import "zookeepergo/network"

//Leader perform operations a leader is supposed to do
func Leader() {
	//connect to client

	//load znode

	//sync with follower

	//deal with message by select
	/*select {
	case Message := <-cR:

	case Message := <-cP
	}*/
	for {

	}

}

//Follower perform operations a follower is supposed to do
func Follower(cR chan network.NetMessage, Peerset []network.Peer, winner int) {
	//load znode

	//sync with leader

	//deal with message by select
	for {
		select {
		case Message := <-cR:
			if Message.Type == 4 && Peerset[Message.Info].Sid == winner {
				return
			}
		}
	}
}
