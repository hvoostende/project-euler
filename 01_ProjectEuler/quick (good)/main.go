package main

import "fmt"

func totalNumbers(below int) int {
	var total int
	for i := 1; i < below; i++ {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	return total
}

func main() {
	fmt.Println(totalNumbers(1000))
}
