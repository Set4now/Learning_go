package main

import "fmt"

func main() {
	mymap := make(map[string]string)
	mymap["deb_suman"] = "devops"
	mymap["das_sunetra"] = "Trips"
	mymap["deb_leena"] = "Housework"

	// fmt.Printf("%v", mymap)
	// for i := 0; i < len(mymap); i++ {

	for key, value := range mymap {
		// fmt.Printf("%V\t\n", index)

		fmt.Printf("%v-->%v\n", key, value)
	}

}
