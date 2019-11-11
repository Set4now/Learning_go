package main

import "fmt"

func main() {
	// anonymous structs
	s := struct {
		name                     string
		emmergencyContacts       map[string]int
		RelationshopWithContacts map[string]string
	}{
		name: "Suman Deb",
		emmergencyContacts: map[string]int{
			"Gopal Deb": 12345677,
			"Leena Deb": 34978346,
		},
		RelationshopWithContacts: map[string]string{
			"Gopal Deb": "Father",
			"Leena Deb": "Mother",
		},
	}

	// fmt.Println(s)
	fmt.Printf("The name of the person is %v\n", s.name)
	fmt.Printf("The emergeny contact details for this person :- \n")
	for k, v := range s.emmergencyContacts {
		fmt.Printf("%v %v\n", k, v)
	}
	fmt.Printf("The person's relationship with the emergency contacts are :- \n")
	for k, v := range s.RelationshopWithContacts {
		println(k, v)
	}
}
