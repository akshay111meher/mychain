package peer

import(
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	"fmt"
	. "../models"
)
var bc Blockchain
var origin = "http://mychain0/"
func StartPeer(){
	fmt.Println("peer started")
	bc = LoadBlockchain()
	http.Handle("/addPeer",websocket.Handler(peerHandler))
	http.Handle("/getBlock", websocket.Handler(blockHandler))
	http.Handle("/sendBlock",websocket.Handler(sendBlockHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}

func sendBlockHandler(ws *websocket.Conn){
	var b Block
	err := websocket.JSON.Receive(ws,&b);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(bc.AddBlock(b));
	bc.SaveChainUsingRoot()
	// bc.CheckAdditionalBlocks()
	if ws.RemoteAddr().String() == origin{
		err = websocket.JSON.Send(ws,Request{100,"success"})
		
		if err != nil {
			log.Fatal(err)
		}	
	}else{
		err = websocket.JSON.Send(ws,Request{100,"stopMining"})
		
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("LatestBlock:",bc.GetLatestBlock().Index, bc.GetLatestBlock().Hash)
}
func peerHandler (ws *websocket.Conn){
	var r Request
	err := websocket.JSON.Receive(ws,&r)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Receive: ", r)
	block := bc.GetLatestBlock()
	err = websocket.JSON.Send(ws,block)
	
	if err != nil {
		log.Fatal(err)
	}
}

func blockHandler(ws *websocket.Conn){
	var r Request
	err := websocket.JSON.Receive(ws,&r)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Receive: ", r)
	block := bc.GetNthBlockFromRoot(r.Number)
	err = websocket.JSON.Send(ws,block)
	
	if err != nil {
		log.Fatal(err)
	}

}
