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
const SV_LISTEN_ADDRESS = "192.168.1.84:63955"

func ServerTransmitter(sendchan chan ConfigFile.ResponseStruct){
	//println("Transmitter Started...SERVER")
	fmt.Printf("bruker fmt \n")
	for{
		select{
		case SendStruct := <- sendchan:
//			fmt.Printf("sensed something to send\n")
//			fmt.Printf("struct ser slik ut: %+v \n", SendStruct)
			arg, _ := json.Marshal(SendStruct)
//			fmt.Printf("sending to %+v", SendStruct.Recipient)
//			fmt.Printf("sending the struct%+v", SendStruct.Response)
			SendStruct.Recipient.Write(arg)
	}
	}
}

func SendHistory(mappet ConfigFile.HistoryStruct,conn *net.TCPConn ){
	arg, _ := json.Marshal(mappet)
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

func ClientListener(conn *net.TCPConn, RecieveChan chan ConfigFile.Request){
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
	var NewReq ConfigFile.Request
	json.Unmarshal(buf[:n], &NewReq)
	NewReq.Adress = conn.RemoteAddr().String()
	NewReq.Socket = conn

	RecieveChan <- NewReq
	}
}
	
func FromServerListener(conn *net.TCPConn, RecieveChan chan ConfigFile.ResponseStruct){
	for{
		buf := make([]byte, 1024)
		n, _ := conn.Read(buf)		//difference read og readfrom??
	//	if (err != nil){println("ERROR i FromServerListener")}
		var NewReq ConfigFile.ResponseStruct
		json.Unmarshal(buf[:n], &NewReq)
		RecieveChan <- NewReq
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
    }

    for {
        conn, err := net.DialTCP("tcp", nil, remote)
        if err == nil {
            return conn
        }
        log.Println("Could not connect to server. Retrying...")
        time.Sleep(1 * time.Second)
    }
}