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

	persons := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	for _, v := range persons {
		fmt.Println(v.firstName)
		fmt.Println(v.lastName)
		for i, val := range v.favouriteIceCreams {
			fmt.Println(i, val)
		}
		fmt.Println("------------")
	}
}
