// This is a server for KTN written in GOlang!

package main

import (
	. "fmt"
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
	Type MsgType
	Msgstring string
	Username string
	IP string
}

type Client struct{
	IP string
	Username string
}

var AllClients map[string]*Client

func main() {
	Printf("Server running... \n")
	msgTx := make(chan Message)
	// start NW sender med msgTx
	msgRx := make(chan Message)
	// start NW mottaker med msgRx

	for {
		select{
		case recieved := <- msgRx:
			Printf("the server has recieved a msg! \n")
			switch(recieved.Type){
				case Login:
					// add IP/username to list #DONE#
					temp := Client{}
					temp.Username = recieved.Username
					temp.IP = recieved.IP
					AllClients[temp.IP] = &temp

				case Logout:
					// remove username from list #DONE#
					delete(AllClients, recieved.IP)

				case Msg:
					temp := Message{}
					temp.Type = Msg
					temp.Username = recieved.Username
					temp.Msgstring = recieved.Msgstring
					for keys := range AllClients{
						// send msg! -legge ved IP?
						Printf("key:  %d\n", keys)

					}

				case Names:
					// send all names back
					for keys := range AllClients{
						// send msg!
						Printf("key:  %d\n", keys)
					}

				case Help:
					// send predetermined info back #DONE#
					temp := Message{}
					temp.Username = "Server: \n"
					temp.Msgstring = "login <username> - log in with the given username\n logout - log out \n msg <message> - send message \n names - list users in chat \n help - view help text \n"
					msgTx <- temp

			}
		default:
			continue
		}

	}
}