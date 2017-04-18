/*
2520 is the smallest number that can be divided by each of the numbers from
1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the
numbers from 1 to 20?
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	counter := 1
outer:
	for {
		maxNumber := 20
		succeses := 0
		for i := 1; i < maxNumber; i++ {
			if counter%i == 0 {
				succeses++
				if succeses == maxNumber-1 {
					fmt.Println(counter)
					fmt.Println(time.Since(startTime))
					break outer
				}
			} else {
				break
			}
		}
		counter++
	}
}
