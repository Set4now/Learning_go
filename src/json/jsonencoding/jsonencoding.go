package main

import (
	"encoding/json"
	"os"
)

// User is a custom struct
type User struct {
	Name string
	Age  int
}

func main() {
	u1 := User{
		Name: "Suman Deb",
		Age:  32,
	}
	u2 := User{
		Name: "Sunetra Das",
		Age:  26,
	}
	userss := []User{u1, u2}
	// json.NewEncoder returns a new encoder that writes to os.Stdout here
	// Encode writes the JSON encoding of v to the stream, followed by a newline character also returns error if any
	err := json.NewEncoder(os.Stdout).Encode(userss)
	if err != nil {
		println("Something wrong", err)
	}

	// println(data)

}
