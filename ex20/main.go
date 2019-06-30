package main

import "fmt"

func main() {
	// Switch with variable will match the value with the cases
	favSport := "ESport"
	switch favSport {
	case "Cricket":
		fmt.Println("Should not print")
	case "Basketball":
		fmt.Println("Basketball")
		// The following should not be executed as
		// there's no "fallthrough" by default in go
		// To enable fallthrough after a particular case, add the
		// keyword fallthrough
	// can also match multiple cases
	case "Golf", "Soccer", "Racing":
		fmt.Println("Others")
	case "ESport":
		fmt.Println("CS:GO")
	}
}
