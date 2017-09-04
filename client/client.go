package main

import(
	"log"
	"golang.org/x/net/websocket"
	"fmt"
)

var origin = "http://localhost/"
var startMiner = "ws://10.200.208.52:8080/startMining"

func main(){
	MinerSignal(5)
}

func MinerSignal(n int){
	for i:=0;i<n;i++{
		sendRequest(startMiner)
	}
}
func sendRequest(url string){
	
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	
	err = websocket.JSON.Send(ws,"")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("sent data")
}