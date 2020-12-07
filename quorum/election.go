package quorum

import (
	"fmt"
	"math/rand"
	"net"
	"time"
	"zookeepergo/network"
)

//ListenCount count down to listen to others
func ListenCount(c chan int) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	c <- 1
}

//ListenCountTally count down to elect oneself as leader
func ListenCountTally(c chan int) {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	c <- 1
}

//LookForLeader either listen to other node or count to elect oneself
func LookForLeader(Peerset []network.Peer, Sid int, Conn []*net.Conn, Response []*net.Conn, cR chan network.NetMessage) int {
	rand.Seed(int64(Sid) + time.Now().UnixNano())
	xid := rand.Intn(100000)
	//choose oneself as leader as initialization
	var vote []int
	network.Winner = -1
	nowVote := -1
	cC := make(chan int)
	cCt := make(chan int)
	go ListenCount(cC)
	//nowState=1 : listening to others
	//nowState=2 : electing myself
	//nowState=3 : winner is decided
	nowState := 1
	for {
		select {
		case <-cC:
			if nowState == 1 {
				nowState = 2
				nowVote = Sid
				for i := 0; i < len(Conn); i++ {
					network.SendMessage(Conn[i], Sid, 1, xid)
				}
				fmt.Println("begin elect myself")
				vote = []int{}
				go ListenCountTally(cCt)
			}
		case <-cCt:
			if nowState == 2 {
				var tally [3]int
				tally[nowVote] = tally[nowVote] + 1
				for i := 0; i < len(vote); i++ {
					tally[vote[i]] = tally[vote[i]] + 1
					//fmt.Printf("*%d ", tally[vote[i]])
				}
				for i := 0; i < len(tally); i++ {
					if tally[i] >= 2 {
						nowVote = i
						network.Winner = i
					}
					fmt.Printf("%d：%d ", i, tally[i])
				}
				fmt.Printf("\n")
				if network.Winner == -1 {
					fmt.Printf("still looking for winner\n")
					nowState = 1
					fmt.Println("begin count down")
					go ListenCount(cC)
					continue
				}
				fmt.Printf("winner is %d\n", network.Winner)
				if network.Winner == Sid {
					//become leader
					fmt.Println("i become leader")
					for i := 0; i < len(Conn); i++ {
						network.SendMessage(Conn[i], Sid, 3, Sid)
					}
					nowState = 3
					return network.Winner
				} /*else {
					//become follower
					fmt.Println("i become follower")
					nowState = 3
					return winner
				}*/
			} else {
				nowState = 1
				fmt.Println("begin count down")
				nowVote = -1
				go ListenCount(cC)
			}
		case Message := <-cR:
			if Message.Type == 1 {
				if nowState == 1 {
					if nowVote == -1 {
						nowVote = Peerset[Message.Id].Sid //translate from relative id to absolute id
					}
					fmt.Printf("vote for %d\n", nowVote)
					network.SendMessage(Conn[Message.Id], Sid, 2, nowVote)
				} else if nowState == 2 {
					fmt.Printf("vote for %d\n", nowVote)
					network.SendMessage(Conn[Message.Id], Sid, 2, nowVote)
				} else if nowState == 3 {
					fmt.Printf("vote for %d\n", network.Winner)
					network.SendMessage(Conn[Message.Id], Sid, 2, network.Winner)
				}
			} else if Message.Type == 2 {
				if nowState == 2 {
					vote = append(vote, Message.Info)
				}

			} else if Message.Type == 3 {
				nowState = 3
				nowVote = Message.Info
				network.Winner = Message.Info
				fmt.Println("i become follower")
				return network.Winner
			}
		}
	}

}

/*
//below is a simple version of leader election used in zookeeper-3.0.0
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
