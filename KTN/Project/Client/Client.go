// This is a server for KTN written in GOlang!

package main

import (
	"fmt"
	"../ConfigFile"
//	"net"
	. "../Network"
	"bufio"
	"os"
	"strings"
)


func main(){
	// start NN
	//
	fmt.Printf("Lift off \n \n")
	IOInputChan := make(chan string)
	RecieveChan := make(chan ConfigFile.ResponseStruct)
	SendChan := make(chan ConfigFile.Request)
	server_addr := "192.168.1.84:63955"

	conn := ConnectToServer(server_addr)

	fmt.Printf("conn er oppe med : %+v \n", conn)

	go FromServerListener(conn, RecieveChan)
	go userInnput(IOInputChan)
	go ClientTransmitter(SendChan, conn)

	for{
		select{
		case NewInput := <- IOInputChan:
		//	println("new NewInput")
		//	println(NewInput)
			temp := makeRequestStruct(NewInput)
		//	fmt.Printf("strucket som sendes er %+v \n", temp)
			if temp.Request != ConfigFile.ERROR{
			//	fmt.Printf("struct is: %+v \n", temp)
				SendChan <- temp
			}else{
				fmt.Printf("you dumb ass motherfucker.....")
			}
		case Respose := <- RecieveChan:
			println(Respose.Timestamp, " ", Respose.Sender," : ", Respose.Content)
		}			

	}

}

func userInnput(IOInputChan chan string){
	reader := bufio.NewReader(os.Stdin)
	for{
		text, _ := reader.ReadString('\n')
		IOInputChan <- text
	}
}

func makeRequestStruct(text string) ConfigFile.Request{
	temp := strings.SplitN(text, " ", 2)
	var returnStruct ConfigFile.Request
	switch temp[0]{
	case ConfigFile.LOGIN:
		returnStruct.Request = ConfigFile.LOGIN
		returnStruct.Content  = strings.TrimPrefix(temp[1], " ")
		lengden := len(returnStruct.Content)
		returnStruct.Content = returnStruct.Content[:lengden-2]
	case ConfigFile.LOGOUT:
		returnStruct.Request = ConfigFile.LOGOUT
	case ConfigFile.MSG:
		returnStruct.Request = ConfigFile.MSG
		returnStruct.Content = strings.TrimPrefix(text, ConfigFile.MSG)
	case ConfigFile.NAMES:
		returnStruct.Request = ConfigFile.NAMES
	case ConfigFile.HELP:
		returnStruct.Request = ConfigFile.HELP
	default:
		fmt.Printf("\n you typed command: %+v \n", temp[0])
		returnStruct.Request = ConfigFile.ERROR
	}
	return returnStruct
}