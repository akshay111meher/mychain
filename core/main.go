package main

import (
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	"strconv"
	"time"
	"log"
	"net/http"
	"golang.org/x/net/websocket"
	)

type Blockchain struct{
	Blocks []Block
}
type Block struct{
	Index string
	PreviousHash string
	Timestamp string
	Data string
	Hash string
}

var bc Blockchain;
type CalculateHash interface{
	SHA256() []byte
}

type BlockChainFunctions interface{
	GetLatestBlock() Block
	GenerateNextBlock(blockData string) Block
	IsValidNewBlock(newBlock Block) bool
	AddBlock(newBlock Block) bool
	IsValidChain() bool
}

type Request struct{
	Number int
	Data string
}

func main() {
	fmt.Println("programme started");
	fmt.Println("blocks start for 0 to n-1");
	var blockChain []Block;

	bc = Blockchain{blockChain}

	//Adding 6 blocks to the chain
	bc.Blocks = append(bc.Blocks,getGenesisBlock());
	secondBlock := bc.GenerateNextBlock("This is second blockdata")
	bc.AddBlock(secondBlock);

	thirdBlock := bc.GenerateNextBlock("This is third blockdata")
	bc.AddBlock(thirdBlock);

	fourthBlock := bc.GenerateNextBlock("This is fourth blockdata")
	bc.AddBlock(fourthBlock);

	fifthBlock := bc.GenerateNextBlock("This is fifth blockdata")
	bc.AddBlock(fifthBlock);

	sixthBlock := bc.GenerateNextBlock("This is sixth blockdata")
	bc.AddBlock(sixthBlock);
	//adding 6 blocks complete

	//generating 7th block
	seventhBlock := bc.GenerateNextBlock("This is seventh blockdata")
	//let us modify hash of seventh block and add this to chain
	//uncomment the below and try to add
	// seventhBlock.hash = "asjkfhskjdhkjasdhk"
	bc.AddBlock(seventhBlock);


	eightBlock := bc.GenerateNextBlock("This is eight blockdata. My name is akshay")
	bc.AddBlock(eightBlock);


	// fmt.Println(bc.IsValidChain());

	// fmt.Println(bc.Blocks)

	http.Handle("/addPeer", websocket.Handler(peerHandler))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

func (b Block) SHA256() string{
	byteArray := sha256.Sum256([]byte(b.Index+b.PreviousHash+b.Timestamp+b.Data))
	return hex.EncodeToString(byteArray[:]);
}

func getGenesisBlock() Block{
	b := Block{"0","0","20170823181145","this is genesis block",""};
	hashByte:= b.SHA256();
	b.Hash = string(hashByte[:])
	return b;
}

func (bc *Blockchain) GetLatestBlock() Block{
	return bc.Blocks[len(bc.Blocks) - 1];
}

func (bc *Blockchain) GetNthBlock(n int) Block{
	if(n >= len(bc.Blocks)){
		return bc.GetLatestBlock()
	}else if(n<=0){
		return getGenesisBlock()
	}else{
		return bc.Blocks[n]
	}
}
 
func (bc *Blockchain) GenerateNextBlock(blockData string) Block{
	previousBlock:= bc.GetLatestBlock();
	currentIndex,_ := strconv.Atoi(previousBlock.Index);
	nextIndexNum:= currentIndex+1;
	nextIndex := strconv.Itoa(nextIndexNum)
	nextTimeStamp  := time.Now().Format("20060102150405")
	nextBlock := Block{nextIndex,previousBlock.Hash,nextTimeStamp,blockData,""};
	hashByte:= nextBlock.SHA256();
	nextBlock.Hash = string(hashByte[:])
	return nextBlock
}

func (bc *Blockchain) IsValidNewBlock(newBlock Block) bool {
	
	latestBlockIndex,_:= strconv.Atoi(bc.GetLatestBlock().Index);
	newBlockIndex,_ := strconv.Atoi(newBlock.Index);
	if(newBlockIndex != latestBlockIndex+1){
		fmt.Println("invalid index")
		return false
	}else if(bc.GetLatestBlock().Hash != newBlock.PreviousHash){
		fmt.Println("invalid previous hash")
		return false
	}else if(newBlock.Hash != newBlock.SHA256()){
		fmt.Println("hashes computed dont match")
		return false
	}
	return true
}

func (bc *Blockchain) AddBlock(newBlock Block) bool{
	if(bc.IsValidNewBlock(newBlock)){
		bc.Blocks = append(bc.Blocks,newBlock)
		fmt.Println("new block "+newBlock.Index+" appended to chain")
		return true
	}
	fmt.Println("new block "+newBlock.Index+" rejected from chain")
	return false
}

func (bc *Blockchain) IsValidChain() bool{
	for i := 1; i < len(bc.Blocks); i++ {
		if(bc.Blocks[i-1].IsNextBlockValid(bc.Blocks[i])){
			continue;
		}else{
			return false;
		}
	}
	return true;
}

func (b Block) IsNextBlockValid(nextBlock Block) bool{
	currentIndex,_ := strconv.Atoi(b.Index);
	newBlockIndex,_ := strconv.Atoi(nextBlock.Index);
	if(currentIndex + 1 != newBlockIndex){
		fmt.Println("Index mismatch")
		return false;
	}else if(b.Hash != nextBlock.PreviousHash){
		fmt.Println("Hash mismatch with previous block")
		return false;
	}else if(nextBlock.Hash != nextBlock.SHA256()){
		fmt.Println("Current block hash computed wrongly")
		return false
	}else{
		return true
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