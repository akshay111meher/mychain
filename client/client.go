package main

import (
	"fmt"
	"log"
	. "../models"
	"golang.org/x/net/websocket"
	. "../controller"
	"encoding/json"
	. "../crypto"
	"math/rand"
)


var origin = "http://localhost/"
var urlGetBlock = "ws://10.200.208.52:8081/getBlock"
var urlSendBlock = "ws://10.200.208.52:8081/sendBlock"
var ownIpSendBlock = "ws://10.200.208.52:8080/sendBlock"
var addPeerIp = "ws://10.200.208.52:8081/addPeer"
var bc Blockchain
func main() {

	bc = LoadBlockchain();
	// bc.PrintChainFromRoot()
	generateBlock(1)
	// sendRequest(urlGetBlock,Request{2,"block"})
	bc.SaveChainUsingRoot();
}

func sendRequest(url string, message Request){
	
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	
	err = websocket.JSON.Send(ws,message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Send:", message)

	var data Block
	err = websocket.JSON.Receive(ws, &data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", data)
}

func sendBlock(url string, block Block){
	
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}
	
	err = websocket.JSON.Send(ws,block)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Index,block.Hash)
	// fmt.Println("Send:", block)

	// var data Block
	// err = websocket.JSON.Receive(ws, &data)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Receive: %s\n", data)
}

func generateBlock(n int){
	for i:=0;i<n;i++ {
		privKey,pubKey := GetAccount("akshay")
		value := randSeq(40)
		r,s := GetSignature(value,privKey)
		var d = Data{value,pubKey,r,s}
		bytes,_ := json.Marshal(d)

		nextBlock := bc.GenerateNextBlock(string(bytes))
		// sendBlock(urlSendBlock,nextBlock)
				//uncomment this to add it to your own chain
		sendBlock(ownIpSendBlock,nextBlock)
		//This will send the block to the connected peers
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}