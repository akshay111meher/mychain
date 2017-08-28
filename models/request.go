package models

import(
	
)

type Request struct{
	Number int
	Data string
}

func Test() Request{
	return Request{1,"test"}
}