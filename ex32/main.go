package main

import "fmt"

type person struct {
	firstName          string
	lastName           string
	favouriteIceCreams []string
}

func main() {
	p1 := person{
		firstName:          "Abhyudit",
		lastName:           "Jain",
		favouriteIceCreams: []string{"Butterscotch", "Vanilla"},
	}

	p2 := person{
		firstName:          "Tutu",
		lastName:           "Bhaiyya",
		favouriteIceCreams: []string{"Chocolate", "Strawberry"},
	}

	fmt.Println(p1.firstName)
	fmt.Println(p1.lastName)
	for i, v := range p1.favouriteIceCreams {
		fmt.Println(i, v)
	}

	fmt.Println(p2.firstName)
	fmt.Println(p2.lastName)
	for i, v := range p2.favouriteIceCreams {
		fmt.Println(i, v)
	}
}
