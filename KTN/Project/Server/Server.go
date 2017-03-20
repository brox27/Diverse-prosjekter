// This is a server for KTN written in GOlang!

package main

import (
	"fmt"
	"../ConfigFile"
	"net"
	. "../Network"
	"time"
	"unicode"
)


func main(){
	// start NN
	//
	fmt.Printf("Lift off Server \n \n")

	RecieveChan := make(chan ConfigFile.Request)
	SendChan := make(chan ConfigFile.ResponseStruct)
	NewConnectionChan := make(chan *net.TCPConn)
	AllClients :=  make(map[string]*ConfigFile.Client)

	go ConnectionListener(NewConnectionChan)
	go ServerTransmitter(SendChan)

	for{
		select{
			case NewMsg := <- RecieveChan:
				fmt.Printf("struct i recieved: \n %+v \n\n", NewMsg)
				var response ConfigFile.ResponseStruct
				response.Recipient = NewMsg.Socket
				response.Sender = "Server"
				response.Timestamp = getTime()

				switch NewMsg.Request{
				case ConfigFile.LOGIN:
					if (isValidLogin(NewMsg, AllClients)){
						fmt.Printf("sensed a login with Username %+v \n", NewMsg.Content)
						AllClients[NewMsg.Adress].Username = NewMsg.Content
						response.Content = "Login Successful"
						SendChan <- response
						// SEND HISTORY

					}else{

						response.Content = "Login Error"
						SendChan <- response
					}
				case ConfigFile.LOGOUT:
					if isLoggedIn(NewMsg, AllClients){
						response.Content = "Logout Successful"
						SendChan <- response
						// Close TCP connection
						AllClients[NewMsg.Adress].Socket.Close()
						delete(AllClients, NewMsg.Adress)
					}else{
						response.Content = "Logout ERROR"
						SendChan <- response
					}
				case ConfigFile.MSG:
					if isLoggedIn(NewMsg, AllClients){
						for key := range AllClients{
							response.Sender = AllClients[NewMsg.Adress].Username
							response.Content = NewMsg.Content
							response.Recipient = AllClients[key].Socket
							SendChan <- response
						}
					}else{
						response.Content = "ERROR"
						SendChan <- response
					}
				case ConfigFile.NAMES:
					response.Content = "Logged in Users are: \n"
					for key := range AllClients{
						response.Content = response.Content + AllClients[key].Username + "\n"
					}
					SendChan <- response

				case ConfigFile.HELP:
					response.Content = "good luck.. cause I wont help you..: \n"
					SendChan <- response
				}
			case NewConnection := <- NewConnectionChan:
					address := NewConnection.RemoteAddr().String()
					var temp ConfigFile.Client
					temp.Username = ""
					temp.Socket = NewConnection
					fmt.Printf("socket address jeg sender er: %+v \n", NewConnection)
					AllClients[address] = &temp
					go ClientListener(NewConnection, RecieveChan)
					fmt.Printf("new Connection from %+v \n", address)
		}
		}
	for key := range AllClients{
	AllClients[key].Socket.Close()
	}

}


func isValidLogin(msg ConfigFile.Request, AllClients map[string]*ConfigFile.Client) bool{
	flag := true
	if !isLoggedIn(msg, AllClients){
		for key := range AllClients{
			if AllClients[key].Username == msg.Content{
				flag = false
			}
		}

	numerals := unicode.Range16{48, 57, 1}
    upper_a_z := unicode.Range16{65, 90, 1}
    lower_a_z := unicode.Range16{97, 122, 1}
    var ranges unicode.RangeTable
    ranges.R16 = []unicode.Range16{numerals, upper_a_z, lower_a_z}
    ranges.LatinOffset = 10 + 26 + 26

    for _, rune := range msg.Content {
    	fmt.Printf("%s \n", rune)
        if !unicode.In(rune, &ranges) {
        	fmt.Printf("** \n")
            flag = false
        }
    }
	}else{
		flag = false
	}
	return flag
}

func isLoggedIn(msg ConfigFile.Request, AllClients map[string]*ConfigFile.Client) bool{
	if AllClients[msg.Adress].Username != ""{
		return true
	}else{
		return false
	}
}

func getTime() string{
	const setup = "Jan 2 15:04"
    return time.Now().Format(setup)
}