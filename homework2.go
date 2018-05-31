package main

import (
	"fmt"
)

func main() {

	pTriple(100)

}

func pTriple(n int) {

	for a := 1; a < n; a++ {
		for b := 1; b < n; b++ {
			for c := 1; c < n; c++ {
				if c*c == a*a+b*b {
					fmt.Println(a, " ", b, " ", c)
				}
			}
		}
	}

}
