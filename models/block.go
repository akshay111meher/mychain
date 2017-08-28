package models

import(
	"strconv"
	"fmt"
	. "../crypto"
)


type Block struct{
	Index string
	PreviousHash string
	Timestamp string
	Data string
	Hash string
}

func (b Block) SHA256() string{
	return SHA256(b.Index+b.PreviousHash+b.Timestamp+b.Data)
}

func (b Block) IsNextBlockValid(nextBlock Block) bool{
	currentIndex,_ := strconv.Atoi(b.Index);
	newBlockIndex,_ := strconv.Atoi(nextBlock.Index);
	if(currentIndex + 1 != newBlockIndex){
		fmt.Println("Index mismatch. Index "+nextBlock.Index)
		return false;
	}else if(b.Hash != nextBlock.PreviousHash){
		fmt.Println("Hash mismatch with previous block. Index "+nextBlock.Index)
		return false;
	}else if(nextBlock.Hash != nextBlock.SHA256()){
		fmt.Println("Current block hash computed wrongly. Index "+nextBlock.Index)
		return false
	}else{
		return true
	}
}