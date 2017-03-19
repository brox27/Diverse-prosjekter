// This is a ConfigFIle for KTN written in GOlang!

package ConfigFile

import (
"net"
)

type ResponseStruct	 struct{
	Recipient *net.TCPConn	`json:"-"`
	Timestamp string	`json:"timestamp"`
	Sender string		`json:"sender"`
	Response string 	`json:"response"`// might be JSON struct
	Content string		`json:"content"`
}

type Request struct{
	Socket *net.TCPConn	`json:"-"`// must be changed
	Adress string	`json:"-"`// must be changed
	Request string	`json:"request"`	//
	Content string	`json:"content"`
}

const (
	LOGIN = "login"
	LOGOUT = "logout"
	MSG = "msg"
	NAMES = "names"
	HELP = "help"
	ERROR = "error"
)

type Client struct{
	Username string
	Socket *net.TCPConn	// må endres
}