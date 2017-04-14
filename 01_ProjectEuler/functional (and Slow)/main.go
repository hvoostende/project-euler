package main

import (
	"fmt"
	"sort"
)

/*
If we list all the natural numbers below 10 that are multiples of 3 or 5,
we get 3, 5, 6 and 9. The sum of these multiples is 23.
Find the sum of all the multiples of 3 or 5 below 1000.
*/

func isMultipleOf(number, multipleOf int) bool {
	if number != 0 {
		if number%multipleOf == 0 {
			return true
		}
	}
	return false
}

func isInSliceInt(item int, list []int) bool {
	for _, v := range list {
		if item == v {
			return true
		}
	}
	return false
}

func multiplesOfNumberBelow(below int, number ...int) []int {
	var listOfNumbersBelow []int
	for _, v := range number {
		for index := 0; index < below; index++ {
			if isMultipleOf(index, v) {
				if !isInSliceInt(index, listOfNumbersBelow) {
					listOfNumbersBelow = append(listOfNumbersBelow, index)
				}
			}
		}
	}
	sort.Ints(listOfNumbersBelow)
	return listOfNumbersBelow
}

func addNumbers(number ...int) int {
	var total int
	for _, v := range number {
		total += v
	}
	return total
}

func main() {
	fmt.Println(addNumbers(multiplesOfNumberBelow(1000, 3, 5)...))
}
