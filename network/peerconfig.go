package network

//Peer save peer information
type Peer struct {
	Sid          int
	Addr         string
	Port         int
	Portresponse int
}

var Peerset []Peer
var Winner int
