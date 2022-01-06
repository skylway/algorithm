package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	recursiveJSON()
}

type nestedStruct struct {
	Msg  string
	Self *nestedStruct
}

func recursiveJSON() {
	s := nestedStruct{
		Msg: "msg",
	}
	s.Self = &s
	str, _ := json.Marshal(s)
	fmt.Println("fuck", string(str))
}