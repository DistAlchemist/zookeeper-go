package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"zookeepergo/quorum"
)

func main() {
	fmt.Println("server begin")

	fmt.Println("load zoo.cfg")
	fp, err := os.Open("zoo.cfg")
	if err != nil {
		fmt.Println("open failed", err)
		return
	}
	r := bufio.NewReader(fp)
	byte, err := r.ReadBytes('\n')
	if err != nil {
		fmt.Println("read failed", err)
		return
	}
	slice := string(byte)
	slice = strings.Replace(slice, "\n", "", -1)
	slice = strings.Replace(slice, "\r", "", -1)
	Myid, err := strconv.Atoi(slice)
	if err != nil {
		fmt.Println("string to int failed", err)
		return
	}
	fmt.Printf("id of this station is %d \n", Myid)

	fmt.Println("initialize network")

	fmt.Println("load znode")
	fmt.Println("start election")
	quorum.LookForLeader(Myid)
}
