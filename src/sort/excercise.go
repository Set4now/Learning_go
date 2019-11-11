package main

// https://godoc.org/sort first example
import (
	"fmt"
	"sort"
)

//Person ..
type Person struct {
	Name string
	Age  int
}

//ByAge mplements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

//ByName  sort.Interface for []Person based on the Name field.
type ByName []Person

func (a ByName) Len() int           { return len(a) }
func (a ByName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func main() {
	a := []int{100, 200, 10, 17}
	s := []string{"Suman", "suman", "Gopal", "Sunetra", "Leena"}
	sort.Ints(a)    // sort on int
	sort.Strings(s) // sort on strings
	fmt.Println(a)
	fmt.Println(s)

	people := []Person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	sort.Sort(ByAge(people))
	fmt.Println(people)

	sort.Sort(ByName(people))
	fmt.Println(people)

}
