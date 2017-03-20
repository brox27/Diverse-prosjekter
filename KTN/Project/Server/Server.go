// This is a server for KTN written in GOlang!

package main

import (
	"fmt"
	"../ConfigFile"
	"net"
	. "../Network"
)


func main(){
	// start NN
	//
	fmt.Printf("Lift off \n \n")

	RecieveChan := make(chan ConfigFile.Request)
	SendChan := make(chan ConfigFile.ResponseStruct)
	NewConnectionChan := make(chan *net.TCPConn)
	AllClients :=  make(map[string]*ConfigFile.Client)

	go ConnectionListener(NewConnectionChan)

	for{
		select{
			case NewMsg := <- RecieveChan:
				switch NewMsg.Request{
				case ConfigFile.LOGIN:
					if (isValidLogin()){
						AllClients[NewMsg.Adress].Username = NewMsg.Content
						response := new(ConfigFile.ResponseStruct)
						response.Sender = "Server"
						response.Content = "Login Successful"
						SendChan <- *response
						// SEND HISTORY

					}else{
						println("ERROR invalid login")
					}
				case ConfigFile.LOGOUT:
					if isLoggedIn(){
						response := new(ConfigFile.ResponseStruct)
						response.Sender = "Server"
						response.Content = "Logout Successful"
						SendChan <- *response
						// Close TCP connection
						delete(AllClients, NewMsg.Adress)
					}else{
						println("You DUMB fuck...")
					}
				case ConfigFile.MSG:
					if isLoggedIn(){
						for key := range AllClients{
							response := new(ConfigFile.ResponseStruct)
							response.Sender = AllClients[NewMsg.Adress].Username
							response.Content = NewMsg.Content
							response.Recipient = AllClients[key].Socket
							SendChan <- *response
						}
					}else{
						println("You DUMB fuck...")
					}
				case ConfigFile.NAMES:
					response := new(ConfigFile.ResponseStruct)	
					response.Content = "Logged in Users are: \n"
					for key := range AllClients{
						response.Content = response.Content + AllClients[key].Username + "\n"
					}
					response.Sender = "Server"
					response.Recipient = NewMsg.Socket

				case ConfigFile.HELP:
					println("standard svar	")
				}
			case NewConnection := <- NewConnectionChan:
					address := NewConnection.RemoteAddr().String()
					var temp ConfigFile.Client
					temp.Username = ""
					temp.Socket = NewConnection
					AllClients[address] = &temp
					go ClientListener(NewConnection, RecieveChan)
					fmt.Printf("new Connection from %+v \n", address)
		}
		for key := range AllClients{
			 AllClients[key].Socket.Close()
		}
	}
}


func isValidLogin() bool{
	return true
}

func isLoggedIn() bool{
	return true
}