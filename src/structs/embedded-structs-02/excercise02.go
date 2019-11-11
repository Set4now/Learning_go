package main

import "fmt"

// This is an example of an embedded structs

type vehicle struct {
	doors int
	color string
}

func main() {

	type truck struct {
		vehicle   // embedding vehicle declared earlier
		fourwheel bool
	}

	type sedan struct {
		vehicle
		luxary bool
	}

	t1 := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourwheel: true,
	}
	t2 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxary: true,
	}

	fmt.Println(t1)
	fmt.Println(t2)

}
