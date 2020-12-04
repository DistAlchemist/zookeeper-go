package quorum

import (
	"fmt"
	"time"
)

var Myid int

//WaitForLeader counting down the tick to elect oneself as a leader
func WaitForLeader(tick int, c chan int) {
	for i := 1; i <= tick; i++ {
		time.Sleep(1000 * time.Millisecond)
	}
	c <- Myid
}

//ListenToLeader listen to other server for potential leader
func ListenToLeader(c chan int) {

}

//LookForLeader either listen to other node or count to elect oneself
func LookForLeader() int {
	ch := make(chan int, 1)
	var LeaderId int
	//begin initial count down
	tick := 10
	for {
		go WaitForLeader(tick, ch)
		go ListenToLeader(ch)
		LeaderId = <-ch
		close(ch)
		fmt.Println("current choose %d as leader", LeaderId)
		/*if LeaderId == Myid {
			if AskForVote() == true {
				break
			}
		} else {
			if VoteFor(LeaderId) == true {
				break
			}
		}*/
	}
	return LeaderId
}
