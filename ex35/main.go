package main

import "fmt"

func main() {
	person := struct {
		first     string
		last      string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		last:  "Bond",
		friends: map[string]int{
			"Q": 777,
			"M": 888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(person.first)
	fmt.Println(person.last)
	fmt.Println(person.friends)
	fmt.Println(person.favDrinks)

	for k, v := range person.friends {
		fmt.Println(k, v)
	}

	for i, v := range person.favDrinks {
		fmt.Println(i, v)
	}
}
