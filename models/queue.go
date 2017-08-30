package models

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