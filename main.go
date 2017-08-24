package main

import (
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	"strconv"
	"time"
)

type Blockchain struct{
	blocks []Block
}
type Block struct{
	index string
	previousHash string
	timestamp string
	data string
	hash string
}

type CalculateHash interface{
	SHA256() []byte
}

type BlockChainFunctions interface{
	GetLatestBlock() Block
	GenerateNextBlock() Block
	IsValidNewBlock(newBlock Block) bool
	AddBlock(newBlock Block) bool
	IsValidChain() bool
}
func main() {
	fmt.Println("programme started");
	fmt.Println("blocks start for 0 to n-1");
	var blockChain []Block;
	bc := Blockchain{blockChain}

	//Adding 6 blocks to the chain
	bc.blocks = append(bc.blocks,getGenesisBlock());
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


	eightBlock := bc.GenerateNextBlock("This is eight blockdata. May name is akshay")
	bc.AddBlock(eightBlock);


	fmt.Println(bc.IsValidChain());

	// fmt.Println(bc.blocks)
}

func (b Block) SHA256() string{
	byteArray := sha256.Sum256([]byte(b.index+b.previousHash+b.timestamp+b.data))
	return hex.EncodeToString(byteArray[:]);
}

func getGenesisBlock() Block{
	b := Block{"0","0","20170823181145","this is genesis block",""};
	hashByte:= b.SHA256();
	b.hash = string(hashByte[:])
	return b;
}

func (bc *Blockchain) GetLatestBlock() Block{
	return bc.blocks[len(bc.blocks) - 1];
}

func (bc *Blockchain) GenerateNextBlock(blockData string) Block{
	previousBlock:= bc.GetLatestBlock();
	currentIndex,_ := strconv.Atoi(previousBlock.index);
	nextIndexNum:= currentIndex+1;
	nextIndex := strconv.Itoa(nextIndexNum)
	nextTimeStamp  := time.Now().Format("20060102150405")
	nextBlock := Block{nextIndex,previousBlock.hash,nextTimeStamp,blockData,""};
	hashByte:= nextBlock.SHA256();
	nextBlock.hash = string(hashByte[:])
	return nextBlock
}

func (bc *Blockchain) IsValidNewBlock(newBlock Block) bool {
	
	latestBlockIndex,_:= strconv.Atoi(bc.GetLatestBlock().index);
	newBlockIndex,_ := strconv.Atoi(newBlock.index);
	if(newBlockIndex != latestBlockIndex+1){
		fmt.Println("invalid index")
		return false
	}else if(bc.GetLatestBlock().hash != newBlock.previousHash){
		fmt.Println("invalid previous hash")
		return false
	}else if(newBlock.hash != newBlock.SHA256()){
		fmt.Println("hashes computed dont match")
		return false
	}
	return true
}

func (bc *Blockchain) AddBlock(newBlock Block) bool{
	if(bc.IsValidNewBlock(newBlock)){
		bc.blocks = append(bc.blocks,newBlock)
		fmt.Println("new block "+newBlock.index+" appended to chain")
		return true
	}
	fmt.Println("new block "+newBlock.index+" rejected from chain")
	return false
}

func (bc *Blockchain) IsValidChain() bool{
	for i := 1; i < len(bc.blocks); i++ {
		if(bc.blocks[i-1].IsNextBlockValid(bc.blocks[i])){
			continue;
		}else{
			return false;
		}
	}
	return true;
}

func (b Block) IsNextBlockValid(nextBlock Block) bool{
	currentIndex,_ := strconv.Atoi(b.index);
	newBlockIndex,_ := strconv.Atoi(nextBlock.index);
	if(currentIndex + 1 != newBlockIndex){
		fmt.Println("Index mismatch")
		return false;
	}else if(b.hash != nextBlock.previousHash){
		fmt.Println("Hash mismatch with previous block")
		return false;
	}else if(nextBlock.hash != nextBlock.SHA256()){
		fmt.Println("Current block hash computed wrongly")
		return false
	}else{
		return true
	}
}