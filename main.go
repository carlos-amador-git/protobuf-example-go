package main

import (
	"fmt"
	"simple"
)

func main() {
	doSimple()
}

func doSimple() {
	sm := simple.SimpleMessage{
		Id: 12345,
		IsSimple: true,
		Name: "My Simple Message",
		SampleList: []int32{1, 4, 7, 8}, 	
	}

	fmt.Println(sm)
}
