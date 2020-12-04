package main

import (
	"fmt"
)

func (c SendThread) start() {
	fmt.Println("send thread start")
}
func (c EventThread) start() {
	fmt.Println("Event thread start")

}
func main() {
	fmt.Println("client begin")

}
