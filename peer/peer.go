package peer

import(
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
	. "../models"
)
var bc *Blockchain
func StartPeer(blockchain *Blockchain){
	bc = blockchain
	http.Handle("/addPeer", websocket.Handler(peerHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
func peerHandler(ws *websocket.Conn){
	var r Request
	err := websocket.JSON.Receive(ws,&r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Receive: ", r)
	block := bc.GetNthBlock(r.Number)
	err = websocket.JSON.Send(ws,block)
	
	if err != nil {
		log.Fatal(err)
	}

}