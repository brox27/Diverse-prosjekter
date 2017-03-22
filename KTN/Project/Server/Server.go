// This is a server for KTN written in GOlang!

package main

import (
	"fmt"
	"../ConfigFile"
	"net"
	. "../Network"
	"time"
	"unicode"
	"encoding/json"
)


func main(){
	//type AllMsgs []Msges
	AllMsgs := make(map[int] []byte)
		id := 0
	fmt.Printf("Lift off Server \n \n")
	var MsgHistory [] []byte
	var Historien []ConfigFile.ResponseStruct

	RecieveChan := make(chan ConfigFile.ServerRequest)
//	SendChan := make(chan ConfigFile.ResponseStruct)
	NewConnectionChan := make(chan *net.TCPConn)
	AllClients :=  make(map[string]*ConfigFile.Client)

	go ConnectionListener(NewConnectionChan)
//	go ServerTransmitter(SendChan)

	for{
		select{
			case NewMsg := <- RecieveChan:
				fmt.Printf("struct i recieved: \n %+v \n\n", NewMsg)
				var response ConfigFile.ResponseStruct
			//	response.Recipient = NewMsg.Socket
				response.Sender = "Server"
				response.Timestamp = getTime()
				response.Response = ConfigFile.INFO
		//		fmt.Printf("req type er: ", NewMsg.Request)
		//		if(NewMsg.Request == "login"){
		//			println("SCOORE MOTHERFUCKER")c
		//		}

				switch NewMsg.Request{
				case ConfigFile.LOGIN:
					if (isValidLogin(NewMsg, AllClients)){
						fmt.Printf("sensed a login with Username %+v \n", NewMsg.Content)
						AllClients[NewMsg.Adress].Username = NewMsg.Content
						response.Content = "Login Successful"
//						SendChan <- response
						ServerTransmitter2(response, NewMsg.Socket)
						// SEND HISTORY
						var HistoryStruct ConfigFile.HistoryStruct
						HistoryStruct.Timestamp = getTime()
						HistoryStruct.Response = ConfigFile.HISTORY
					//	HistoryStruct.Content = AllMsgs
						HistoryStruct.Content = MsgHistory
					//	Historien2, _ := json.Marshal(Historien)
					//	HistoryStruct.Content = Historien2
						println("*** rett over bror")

				//		HistoryStruct.Content = MsgHistory
						if Historien != nil{
							println("*** SKAL SENDE HISTORY")
							SendHistory(HistoryStruct, NewMsg.Socket)
						}


					}else{
						response.Response = ConfigFile.ERROR
						response.Content = "Login Error"
					//	SendChan <- response
						ServerTransmitter2(response, NewMsg.Socket)
					}
				case ConfigFile.LOGOUT:
					if isLoggedIn(NewMsg, AllClients){
						response.Content = "Logout Successful"
						//SendChan <- response
						ServerTransmitter2(response, NewMsg.Socket)
						AllClients[NewMsg.Adress].Socket.Close()
						delete(AllClients, NewMsg.Adress)
					}else{
						response.Response = ConfigFile.ERROR
						response.Content = "Logout ERROR"
					//	SendChan <- response
						ServerTransmitter2(response, NewMsg.Socket)
					}
				case ConfigFile.MSG:
					if isLoggedIn(NewMsg, AllClients){
					response.Response = ConfigFile.MESSEGE
					response.Sender = AllClients[NewMsg.Adress].Username
					response.Content = NewMsg.Content
					response2, _ := json.Marshal(response)
					MsgHistory = append(MsgHistory, response2)
					Historien = append(Historien, response)
						for key := range AllClients{
							response.Recipient = AllClients[key].Socket
							//println(key)
	//						SendChan <- response
							ServerTransmitter2(response, AllClients[key].Socket)
						}
					arg, _ := json.Marshal(NewMsg)
					AllMsgs[id] = arg
					id++	
					
					}else{
						println("jaaa...")
						response.Response = ConfigFile.ERROR
						ServerTransmitter2(response, NewMsg.Socket)
					}
				case ConfigFile.NAMES:
					response.Content = "Logged in Users are: \n"
					for key := range AllClients{
						response.Content = response.Content + AllClients[key].Username + "\n"
					}
//					SendChan <- response
					ServerTransmitter2(response, NewMsg.Socket)

				case ConfigFile.HELP:
					response.Content = "good luck.. cause I wont help you..: \n"
					ServerTransmitter2(response, NewMsg.Socket)
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
	fmt.Printf("hist: %+v \n", MsgHistory)
	for key := range AllClients{
	AllClients[key].Socket.Close()
	}

}


func isValidLogin(msg ConfigFile.ServerRequest, AllClients map[string]*ConfigFile.Client) bool{
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

func isLoggedIn(msg ConfigFile.ServerRequest, AllClients map[string]*ConfigFile.Client) bool{
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