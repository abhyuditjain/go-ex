package main

import "fmt"

func main() {
	// Switch without any expression/variable is equivalent to switch true
	switch {
	case false:
		fmt.Println("Should not print")
	case 1 == 1:
		fmt.Println("1 == 1")
		// The following should not be executed as
		// there's no "fallthrough" by default in go
		// To enable fallthrough after a particular case, add the
		// keyword fallthrough
	case 2 == 2:
		fmt.Println("2 == 2")
	case 3 == 3:
		fmt.Println("3 == 3")
	}
}
