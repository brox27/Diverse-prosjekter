// This is a server for KTN written in GOlang!

package Network

import (
//	. "fmt"
	"../ConfigFile"
	"encoding/json"
	"time"
	"log"
	"fmt"
	"net"
)
const SV_LISTEN_ADDRESS = "192.168.1.170:12345"
// Raspberry = 

func ServerTransmitter(sendchan chan ConfigFile.ResponseStruct){
	//println("Transmitter Started...SERVER")
	fmt.Printf("bruker fmt \n")
	for{
		select{
		case SendStruct := <- sendchan:
	//		fmt.Printf("sensed something to send\n")
	//\\		fmt.Printf("struct ser slik ut: %+v \n", SendStruct)
			arg, _ := json.Marshal(SendStruct)
//			fmt.Printf("sending to %+v", SendStruct.Recipient)
//			fmt.Printf("sending the struct%+v", SendStruct.Response)
			SendStruct.Recipient.Write(arg)
		}
//
	}
}

func ServerTransmitter2(sendFile ConfigFile.ResponseStruct, conn *net.TCPConn){
	//println("Transmitter Started...SERVER")
//			fmt.Printf("sensed something to send\n")
		println("skal ha sendt1 ")	
//			fmt.Printf("struct ser slik ut: %+v \n", SendStruct)
			arg, _ := json.Marshal(sendFile)
//			fmt.Printf("sending to %+v", SendStruct.Recipient)
//			fmt.Printf("sending the struct%+v", SendStruct.Response)
			conn.Write(arg)
			println("skal ha sendt", string(arg))
	}

func SendHistory(historystruct ConfigFile.HistoryStruct, conn *net.TCPConn ){
//		time.Sleep(2*time.Second)
	fmt.Printf("starter JSON \n")
	arg, _ := json.Marshal(historystruct)
//	time.Sleep(2*time.Second)
	fmt.Printf("sender dette:  \n", arg)
	fmt.Printf("sender stringen:  \n", string(arg))
	conn.Write(arg)
}

func ClientTransmitter(sendchan chan ConfigFile.Request, conn *net.TCPConn){
	println("Transmitter Started...Client")
	for{
	select{
		case SendStruct := <- sendchan:
			arg, _ := json.Marshal(SendStruct)
		//	fmt.Printf("arg: %+v \n", string(arg))
			conn.Write(arg)
	}
	}
}

func ClientListener(conn *net.TCPConn, RecieveChan chan ConfigFile.ServerRequest){
	for{
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)		//difference read og readfrom??
	//	bufff := string(buf)
	//	fmt.Printf("buf er: %+v \n", bufff)
        if err != nil {
            fmt.Printf("feilen er %+v \n", err)
            fmt.Printf("n er ", n)
            return
        }
	var NewReq ConfigFile.ServerRequest
	json.Unmarshal(buf[:n], &NewReq)
	NewReq.Adress = conn.RemoteAddr().String()
	NewReq.Socket = conn

	RecieveChan <- NewReq
	}
}
	
func FromServerListener(conn *net.TCPConn, RecieveChan chan ConfigFile.ResponseStruct, historyChan chan ConfigFile.HistoryStruct){
	for{
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)		//difference read og readfrom??
		if err != nil {
            conn.Close()
            fmt.Println("Lost connection with server")
           // connection_terminated <- true
            return
        }
	//	if (err != nil){println("ERROR i FromServerListener")}
		var NewReq ConfigFile.ResponseStruct
		json.Unmarshal(buf[:n], &NewReq)
		if NewReq.Response == ConfigFile.HISTORY{
			var HistReq ConfigFile.HistoryStruct
			json.Unmarshal(buf[:n], &HistReq)
			historyChan <- HistReq
		}else{
			RecieveChan <- NewReq
		}
		
	}
	conn.Close()
}

func ConnectionListener(NewConnectionChan chan *net.TCPConn){
	println("ConnectionListener Started...")
	local, err := net.ResolveTCPAddr("tcp", SV_LISTEN_ADDRESS)

	ln, err := net.ListenTCP("tcp", local)
	if err != nil {
		fmt.Printf("error in ConnectionListener... could not...")
	}
	for {
		conn, err := ln.AcceptTCP()
		if err != nil {
			fmt.Printf("error in ConnectionListener... could not...")
		}
		NewConnectionChan <- conn
	}
}


func ConnectToServer(addr string) (*net.TCPConn){
    remote, err := net.ResolveTCPAddr("tcp", addr)
    if err != nil {
        log.Fatal(err)
        fmt.Printf("\n*** ERROR i ResolveTCPAddr LES DET UNDER: \n %+v +n", err)
    }

    for {
        conn, err := net.DialTCP("tcp", nil, remote)
        if err == nil {
            return conn
        }
        log.Println("Could not connect to server. Retrying...")
        fmt.Printf("\n LES ALT UNDER:\n %+v \n ", err)
        time.Sleep(1 * time.Second)
    }
}