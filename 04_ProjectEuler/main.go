/*
A palindromic number reads the same both ways. The largest palindrome made from
the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
	"fmt"
	"math"
	"strconv"
)

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func checkPalinDrome(n int) bool {
	s := strconv.Itoa(n) //convert n to string
	if s == reverse(s) {
		return true
	}
	return false
}

func productOfNDigitNumbers(n int) []int {
	var sliceOfProducts []int
	for i := math.Pow(10, float64(n-1)); i <= math.Pow(10, float64(n)); i++ {
		for j := math.Pow(10, float64(n-1)); j <= math.Pow(10, float64(n)); j++ {
			sliceOfProducts = append(sliceOfProducts, int(i)*int(j))
		}
	}
	return sliceOfProducts
}

func main() {
	var largestPalindrome int
	for _, v := range productOfNDigitNumbers(3) {
		if checkPalinDrome(v) {
			if v > largestPalindrome {
				largestPalindrome = v
			}
		}
	}
	fmt.Println(largestPalindrome)
}
