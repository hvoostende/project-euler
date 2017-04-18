package main

import "fmt"

func divisableBy(x, y int) bool {
	if x%y == 0 {
		return true
	}
	return false
}

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			fmt.Println("now im 5")
		} else {
			fmt.Println("not 5")
			fmt.Println()
		}
	}
}
