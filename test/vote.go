package main

import "fmt"

//AskForVote elect oneself as leader and receive other votes
func AskForVote() bool {
	fmt.Println("now ask for vote")
	return true

}

//VoteFor vote for others and wait for result
func VoteFor(LeaderId int) bool {
	fmt.Printf("now vote for %d \n", LeaderId)
	return true
}
