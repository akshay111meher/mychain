package main

import(
	"fmt"
)

type Block struct{
	Index string
	PreviousHash string
	Timestamp string
	Data string
	Hash string
	Nonce string
}

type Blockchain struct{
	Blocks Block
	Next []*Blockchain
	Previous *Blockchain
}

type Queue struct {
	Array []*Blockchain
}

func (q *Queue) Push(bc *Blockchain) bool{
	q.Array = append(q.Array,bc)
	return true;
}
func (q *Queue) Pop () (*Blockchain){
	elem := q.Array[0]
	q.Array = append(q.Array[:0], q.Array[1:]...)
	return elem
}
func (bc *Blockchain) AppendToChain(nextBlock Block) bool{
	if(nextBlock.PreviousHash == bc.Blocks.Hash){
		fmt.Println(bc.Next)
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
	}else{
			return bc.Previous.AppendFromEnd(newBlock)
	}
	return false
}

func main(){
		b := Block{"0","aa","3","4","bb","6"}
		// pointer := new(Blockchain)
		var array []*Blockchain;
		root := Blockchain{b,array,nil}
		// fmt.Println(b)
		b2 := Block{"1","bb","32","42","cc","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))

		b2 = Block{"2","cc","32","42","dd","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))

		b2 = Block{"3","dd","32","42","ee","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))		

		b2 = Block{"4","ee","32","42","hh","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))
				
		b2 = Block{"5","hh","32","42","ii","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))

		b2 = Block{"6","ii","32","42","jj","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))

		b2 = Block{"4","ee","32","42","ff","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))
		
		b2 = Block{"5","ff","32","42","gg","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))

		b2 = Block{"6","gg","32","42","pp","62"}
		// fmt.Println(b2)
		fmt.Println(root.AppendToChain(b2))		
		endNode:= root.GetLatestNode()

		b2 = Block{"7","pp","32","42","qq","62"}
		// fmt.Println(b2)
		fmt.Println(endNode.AppendFromEnd(b2))
		// endNode = root.GetLatestNode()
		fmt.Println(endNode)
}