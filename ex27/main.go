package main

import "fmt"

func main() {
	states := []string{
		`Alabama`,
		`Alaska`,
		`Arizona`,
		`Arkansas`,
		`California`,
		`Colorado`,
		`Connecticut`,
		`Delaware`,
		`Florida`,
		`Georgia`,
		`Hawaii`,
		`Idaho`,
		`Illinois`,
		`Indiana`,
		`Iowa`,
		`Kansas`,
		`Kentucky`,
		`Louisiana`,
		`Maine`,
		`Maryland`,
		`Massachusetts`,
		`Michigan`,
		`Minnesota`,
		`Mississippi`,
		`Missouri`,
		`Montana`,
		`Nebraska`,
		`Nevada`,
		`New Hampshire`,
		`New Jersey`,
		`New Mexico`,
		`New York`,
		`North Carolina`,
		`North Dakota`,
		`Ohio`,
		`Oklahoma`,
		`Oregon`,
		`Pennsylvania`,
		`Rhode Island`,
		`South Carolina`,
		`South Dakota`,
		`Tennessee`,
		`Texas`,
		`Utah`,
		`Vermont`,
		`Virginia`,
		`Washington`,
		`West Virginia`,
		`Wisconsin`,
		`Wyoming`,
	}

	fmt.Println(len(states))
	fmt.Println(cap(states))

	for i := 0; i < len(states); i++ {
		fmt.Println(i, states[i])
	}

	// now to `make` a slice
	newStates := make([]string, 50, 500)
	fmt.Println(len(newStates))
	fmt.Println(cap(newStates))
	// put a value into each of the 50 positions in the length of the slice
	// we are putting values into the 50 positions that are the length of the slice
	for i := 0; i < 50; i++ {
		newStates[i] = fmt.Sprintf("Position %d", i)
	}
	newStates = append(newStates, states...)

	fmt.Println(len(newStates))
	fmt.Println(cap(newStates))

	for i := 0; i < len(newStates); i++ {
		fmt.Println(i, newStates[i])
	}
}
