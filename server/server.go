package main

import (
	"fmt"
)

func main() {
	fmt.Println("server begin")
	fmt.Println("load zoo.cfg")
	fmt.Println("initialize network")
	fmt.Println("load znode")
	fmt.Println("start election")
	LookForLeader()
}
