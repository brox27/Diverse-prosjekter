// This is a server for KTN written in GOlang!

package main

import (
	. "fmt"
	"bufio"
	"os"
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


func main() {

	msgTx := make(chan Message)
	// start NW sender med msgTx
	msgRx := make(chan Message)
	// start NW mottaker med msgRx
	go CheckKeyboard(msgTx)
	Printf("Client running... \n")
	for {
		select{
		case recieved := <- msgRx:
			Printf("the server has recieved a msg! \n")
			switch(recieved.Type){
				case Msg:
					Printf(recieved.Username)
					Printf(": ")
					Printf(recieved.Msgstring)

				default: 
					Printf("ERROR!: unknown message recieved \n")
			}
		default:
			continue
		}

	}
}

func CheckKeyboard(msgTx chan Message){
	temp := Message{}
	loginStatus := false
	for{
		reader := bufio.NewReader(os.Stdin)
	    input, _ := reader.ReadString('\n')
	    if (input[0:5] == "login"){
	    	if loginStatus{
	    		Printf("You are allready logged in... JERK!")
	    		continue
	    	}
	    	temp.Username = input[6:]
	    	temp.Type = Login
	    	loginStatus = true
	    	msgTx <- temp
	    }else if input[0:6] == "logout"{
	    	loginStatus = false
	    	temp.Type = Logout
	    	msgTx <- temp
	    }else if input[0:4] == "help"{
	    	temp.Type = Help
	    	msgTx <- temp
	    }else if input[0:5] == "names"{
	    	temp.Type = Names
	    	msgTx <- temp
	    }else{
	    	temp.Type = Msg
	    	temp.Msgstring = input
	    	msgTx <- temp
	    }

	}
}