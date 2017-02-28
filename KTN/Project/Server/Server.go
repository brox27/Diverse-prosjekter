// This is a server for KTN written in GOlang!

package main

import (
	. "fmt"
	"net"
//	"time"
)
type MsgType int

const (
	Login MsgType = iota
	Logout
	Msg
	Names
	Help
)

type Message struct{
	Timestamp string
	Sender string
	Response string
	Content string
}

type Client struct{
	IP string
	Username string
	Socket *net.TCPConn
}

var AllClients map[string]*Client

func main() {
	Printf("Server running... \n")
	msgTx := make(chan Message)
	conRx := make(chan Message)
	// start NW sender med msgTx
	msgRx := make(chan Message)
	// start NW mottaker med msgRx

	for {
		select{
		case connected := <- conRx:
			thatIP := connected.RemoteAddr().string()
			temp := Client{}
			AllClients[thatIP] = &temp

		case recieved := <- msgRx:
			request := "LOL" // recieved.request
			content := "LOL"// recieved.content
			switch(request){
			case "login":

			case "logout":

			case "msg":

			case "names":

			case "help":

			default:
				Printf("ERROR: unknown request \n")
			}
		default:
			continue
		}

	}
}