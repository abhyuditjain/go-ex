package main

import "fmt"

func main() {
	bd := 1992
	// For without any init; condition; post
	// equivalent to (for ; ;) or (while (true))
	for {
		if bd > 2019 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
