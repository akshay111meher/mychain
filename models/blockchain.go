package models

import(
	"strconv"
	"time"
	"fmt"
	. "../controller"
	. "../consensus"
	"encoding/json"
)


type Blockchain struct{
	Blocks Block
	Next []*Blockchain
	Previous *Blockchain
}

func (root *Blockchain) GetLatestBlock() Block{
	return root.GetLatestNode().Blocks
}

func (root *Blockchain) GetNthBlockFromRoot(n int) Block{
	tail := root.GetLatestNode();
	return tail.GetNthBlock(n)
}

func (tail *Blockchain) GetNthBlock(n int) Block{
	if(tail.Previous == nil){
		nString := strconv.Itoa(n)
		if(nString == tail.Blocks.Index){
			return tail.Blocks
		}else{
		   return Block{}
		}
	}else{
		nString := strconv.Itoa(n)
		if(nString == tail.Blocks.Index){
			return tail.Blocks
		}else{
			return tail.Previous.GetNthBlock(n)
		}
	}
}


func (root *Blockchain) GenerateNextBlock(blockData string) Block{
	previousBlock:= root.GetLatestBlock();
	currentIndex,_ := strconv.Atoi(previousBlock.Index);
	nextIndexNum:= currentIndex+1;
	nextIndex := strconv.Itoa(nextIndexNum)
	nextTimeStamp  := time.Now().Format("20060102150405")
	nextBlock := Block{nextIndex,previousBlock.Hash,nextTimeStamp,blockData,"",""};
	hashByte:= nextBlock.SHA256();
	nextBlock.Hash = string(hashByte[:])
	nextBlock.Nonce = ReturnNonce(nextBlock.Hash)
	return nextBlock
}

func (root *Blockchain) IsValidNewBlock(newBlock Block) bool {

	if(len(root.Next) == 0){
		if(newBlock.IsThisBlockValid()){
			return true
		}else{
		    return false
		}
	}
	latestBlockIndex,_:= strconv.Atoi(root.GetLatestBlock().Index);
	newBlockIndex,_ := strconv.Atoi(newBlock.Index);
	if(newBlockIndex != latestBlockIndex+1){
		fmt.Println("invalid index")
		return false
	}else if(root.GetLatestBlock().Hash != newBlock.PreviousHash){
		fmt.Println("invalid previous hash")
		return false
	}else if(newBlock.Hash != newBlock.SHA256()){
		fmt.Println("hashes computed dont match")
		return false
	}else if(!newBlock.IsThisBlockValid()){
		return false
	}
	return true
}

func (root *Blockchain) SaveChainUsingRoot(){
	tail := root.GetLatestNode();
	tail.SaveChain()
	fmt.Println("Chain saved")
}
func (tail *Blockchain) SaveChain(){
	if tail.Previous == nil{
		return
	}else{
		blockMarshal,_ := json.Marshal(tail.Blocks)
		CreateFile(tail.Blocks.PreviousHash,blockMarshal)
		tail.Previous.SaveChain()
	}
}
func (root *Blockchain) AddBlock(newBlock Block) bool{
	if(root.IsValidNewBlock(newBlock)){
		tail := root.GetLatestNode()
		if tail.AppendFromEnd(newBlock){
			fmt.Println("Received Latest Block ",newBlock.Index)
		}else{
			root.AppendToChain(newBlock)
		}
		fmt.Println("new block "+newBlock.Index+" appended to chain")
		blockMarshal,_ := json.Marshal(newBlock)
		CreateFile(newBlock.PreviousHash,blockMarshal);
		return true
	}
	fmt.Println("new block "+newBlock.Index+" rejected from chain")
	return false
}
func (root *Blockchain) IsValidChainFromEnd() bool{
	tail := root.GetLatestNode()
	return tail.IsValidChain()
}
func (tail *Blockchain) IsValidChain() bool{
	
	if tail.Previous == nil{
		return true
	}else{
		if tail.Previous.Blocks.IsNextBlockValid(tail.Blocks){
			return tail.Previous.IsValidChain()
		}else{
			return false
		}
	}
}
func (root *Blockchain) PrintChainUsingRoot(){
	tail := root.GetLatestNode()
	tail.PrintChain();
}
func (tail *Blockchain) PrintChain(){
	b,err := json.MarshalIndent(tail.Blocks,"","  ")
	if err!= nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
	if tail.Previous == nil{
		return
	}
	tail.Previous.PrintChain()
	
}
func NewBlockchain(data string) (Blockchain){
	
	b := getGenesisBlock(data)
	// pointer := new(Blockchain)
	var array []*Blockchain;
	root := Blockchain{b,array,nil}
	blockMarshal,_ := json.Marshal(b)
	CreateFile(b.PreviousHash,blockMarshal);
	return root
}

func LoadBlockchain() (Blockchain){
	var temp Block
	var array []*Blockchain;
	var previousHash string;
	previousHash = "0"
	
	blockData := ReadFile(previousHash)
	json.Unmarshal(blockData, &temp)
	root := Blockchain{temp,array,nil}
	previousHash = temp.Hash

	for {
		blockData:= ReadFile(previousHash);
		if(len(blockData) == 0){
			break;
		}else{
			json.Unmarshal(blockData,&temp)
			root.AppendToChain(temp)
			previousHash = temp.Hash
		}
	}
	return root
}
func getGenesisBlock(data string) Block{
	b := Block{"0","0","20170823181145",data,"","0"};
	hashByte:= b.SHA256();
	b.Hash = string(hashByte[:])
	return b;
}

func (bc *Blockchain) AppendToChain(nextBlock Block) bool{
	if(nextBlock.PreviousHash == bc.Blocks.Hash){
		newNode := new(Blockchain);
		newNode.Blocks = nextBlock;
		var array []*Blockchain;		
		newNode.Next = array;
		newNode.Previous = bc;
		bc.Next = append(bc.Next,newNode)
		return true
	}else{
		for i:=0;i<len(bc.Next);i++{
			if bc.Next[i].AppendToChain(nextBlock){
				return true
			}
		}
	}
	return false
}

func (bc *Blockchain) GetLatestNode() (*Blockchain){
	var array []*Blockchain
	q := Queue{array}

	q.Push(bc)
	var lastNode *Blockchain
	for;;{
		if(len(q.Array)==0){
			break
		}else{
			lastNode = q.Pop()
			for i:=0;i<len(lastNode.Next);i++{
				q.Push(lastNode.Next[i])
			}
		}	
	}

	return lastNode
}

func (bc *Blockchain) AppendFromEnd(newBlock Block) bool{
	if(bc.Blocks.Hash == newBlock.PreviousHash){
		newNode := new(Blockchain)
		newNode.Blocks = newBlock
		var array []*Blockchain;		
		newNode.Next = array;
		newNode.Previous = bc
		bc.Next = append(bc.Next,newNode)
		return true
	}else{
			if bc.Previous == nil{
				return false
			}else{
				return bc.Previous.AppendFromEnd(newBlock)
			}
	}
}