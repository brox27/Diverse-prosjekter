// This is a server for KTN written in GOlang!

package Nettwork

import (
	. "fmt"
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

func Transmit(msgTx chan Message) {
	// JSON this shit!
	// send on TCP!
}