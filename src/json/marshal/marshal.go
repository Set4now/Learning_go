package main

import (
	"encoding/json"
)

//User ...This is a custom type
type User struct {
	Name string `json:"name"` //change the exported field to this `json:"name"`
	Age  int
}

func main() {
	// u1 := map[string]string{
	// 	"Name":  "suman",
	// 	"skill": "python",
	// }
	// // fmt.Printf("%v", u1["Name"])
	// bs, err := json.Marshal(u1)
	// println(string(bs))
	// println(err)
	// if err != nil {
	// 	println(string(bs))
	// }
	users := []User{}
	u1 := User{
		Name: "SumanDeb",
		Age:  32,
	}
	u2 := User{
		Name: "SunetraDas",
		Age:  26,
	}
	users = append(users, u1, u2)
	bs, _ := json.Marshal(users) // return a slice of bytes []bytes and an optional error
	println(string(bs))
}
