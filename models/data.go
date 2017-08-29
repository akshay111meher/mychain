package models

import(
	"encoding/json"
	"fmt"
)
type Data struct {
	Value string
	PublicKey string
	R string 
	S string
}

func (d Data) ToString() string{
	bytes,_ := json.Marshal(d);
	fmt.Println(string(bytes))
	return string(bytes)
}