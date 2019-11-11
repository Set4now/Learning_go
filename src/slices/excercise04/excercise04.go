package main

import "fmt"

func main() {
	// inislice := []string{`Andhra Pradesh`, `Arunachal Pradesh`, `Assam`, `Bihar`, `Chhattisgarh`}
	var inslice []string
	inslice = []string{`Andhra Pradesh`, `Arunachal Pradesh`, `Assam`, `Bihar`, `Chhattisgarh`}
	fmt.Printf("%v\n", inslice)
	fmt.Printf("%T\n", inslice)
	// fmt.Printf("%v", inislice)
	// mystates := make([]string, 5, 29) // Using make to create a slice of size 5 and capacity 29
	// mystates = append(mystates, inislice...)
	mystates := make([]string, 5, 29) // // Using make to create a slice of size 5 and capacity 29
	// fmt.Printf("%v\n", len(mystates))

	// fmt.Printf("%v\n", cap(mystates))
	mystates = []string{`Andhra Pradesh`, `Arunachal Pradesh`, `Assam`, `Bihar`, `Chhattisgarh`} // adding items
	// OR

	// fmt.Printf("%v\n", len(mystates)) // printing lenth
	// fmt.Printf("%v\n", cap(mystates)) // printing capacity
	// fmt.Printf("%v\n", mystates)
	for i := 0; i < len(mystates); i++ { // printing the items along with their index positions
		fmt.Printf("%v\t%v\n", i, mystates[i])
	}
}
