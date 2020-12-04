package quorum

import (
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"zookeepergo/network"
)

//LookForLeader either listen to other node or count to elect oneself
func LookForLeader(Peerset []network.Peer, Sid int) int {
	var Conn []net.Conn
	Conn = append(Conn, network.ConnectToServer(Peerset[0]))
	Conn = append(Conn, network.ConnectToServer(Peerset[1]))
	xid := rand.Intn(100000)
	//choose oneself as leader as initialization
	var vote []int
	var tally [3]int
	var winner int
	nowVote := Sid
	for i := 0; i < len(Conn); i++ {
		Conn[i].Write([]byte(strconv.Itoa(xid)))
		var readinfo []byte
		Conn[i].Read(readinfo)
		vote1, err := strconv.Atoi(string(readinfo))
		if err != nil {
			fmt.Println("recv error", err)
		}
		vote = append(vote, vote1)
	}
	tally[nowVote] = tally[nowVote] + 1
	tally[vote[0]] = tally[vote[0]] + 1
	tally[vote[1]] = tally[vote[1]] + 1
	for i := 0; i < len(vote); i++ {
		if tally[i] >= 2 {
			nowVote = i
			winner = i
		}
	}
	if winner == Sid {
		//become leader
	} else {
		//become follower
	}
	return nowVote
}

/*
//below is a attampt to write Raft, remain unfinished
//WaitForLeader counting down the tick to elect oneself as a leader
func WaitForLeader(Myid int, tick int, c chan int) {
	for i := 1; i <= tick; i++ {
		time.Sleep(1000 * time.Millisecond)
	}
	c <- Myid
}

//ListenToLeader listen to other server for potential leader
func ListenToLeader(c chan int) {

}

//LookForLeader either listen to other node or count to elect oneself
func LookForLeader(Myid int) int {
	ch := make(chan int, 1)
	var LeaderId int
	Myid = 1
	//begin initial count down
	tick := 10
	for {
		go WaitForLeader(Myid, tick, ch)
		go ListenToLeader(ch)
		LeaderId = <-ch
		fmt.Printf("current choose %d as leader\n", LeaderId)
		if LeaderId == Myid {
			if AskForVote() == true {
				break
			}
		} else {
			if VoteFor(LeaderId) == true {
				break
			}
		}
	}
	close(ch)
	return LeaderId
}*/
