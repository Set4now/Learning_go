package main

import (
	"encoding/json"
	"fmt"
)

type data []struct {
	Name string `json:"name"`
	Age  int    `json:"Age"`
}

// func Unmarshal(data []byte, v interface{}) error

func main() {
	var mydata data // declare a varaible of a type data
	myjson := []byte(`[{"name":"SumanDeb","Age":32},{"name":"SunetraDas","Age":26}]`)
	err := json.Unmarshal(myjson, &mydata) // change the value pointed by this
	if err != nil {
		println("error:", err)
	}
	fmt.Printf("%+v", mydata)
}
