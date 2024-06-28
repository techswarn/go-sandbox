package main

import (
	"encoding/json"
	"fmt"
)

type T struct {
	F1 string `json:"field1"`
	F2 string `json:"field2, omitempty"`
	F3 string `json:"field3, omitempty"`
	F4 string `json:"field4"`
}

func main(){
	t := &T{
		F1:"Same",
		F2:"",
		F3:"Mark",
		F4:"mil",
	}

	v, _ := json.Marshal(t)

	fmt.Printf("Struct representation %+s", v)
}