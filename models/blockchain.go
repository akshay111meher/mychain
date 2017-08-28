package models

import(
	"strconv"
	"time"
	"fmt"
)


type Blockchain struct{
	Blocks []Block
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

func NewBlockchain() (bc Blockchain){
	var blocks []Block
	bc = Blockchain{blocks}
	bc.Blocks = append(bc.Blocks,getGenesisBlock())
	return bc
}
func getGenesisBlock() Block{
	b := Block{"0","0","20170823181145","this is genesis block",""};
	hashByte:= b.SHA256();
	b.Hash = string(hashByte[:])
	return b;
}