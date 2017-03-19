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

	go FromServerListener(RecieveChan)
	go userInnput(IOInputChan)
	go ClientTransmitter(SendChan)

	select{
	case NewInput := <- IOInputChan:
		println("new NewInput")
		println(NewInput)
		temp := makeRequestStruct(NewInput)
		if temp.Request != ConfigFile.ERROR{
			SendChan <- temp
		}else{
			fmt.Printf("you dumb ass motherfucker.....")
		}
	case Respose := <- RecieveChan:
		println("recieved from server")
		println(Respose.Sender)

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
	// TrimPrefix
	switch temp[0]{
	case ConfigFile.LOGIN:
		returnStruct.Request = ConfigFile.LOGIN
		returnStruct.Content = strings.TrimPrefix(text, ConfigFile.LOGIN)
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
		returnStruct.Request = ConfigFile.ERROR
	}
	return returnStruct
}