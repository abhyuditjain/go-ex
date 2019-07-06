package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	truck1 := truck{
		vehicle: vehicle{
			doors:  4,
			colour: "blue",
		},
		fourWheel: true,
	}

	sedan1 := sedan{
		vehicle: vehicle{
			doors:  4,
			colour: "blue",
		},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)
	// If there are no two fields with same name, these will be "promoted"
	fmt.Println(truck1.doors)
	fmt.Println(sedan1.doors)
}
