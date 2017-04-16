/*var
The prime factors of 13195 are 5, 7, 13 and 29.
What is the largest prime factor of the number 600851475143 ?
*/

package main

import (
	"fmt"
	"sort"
)

func maxPrimeFactor(n int) int {
	//init
	largestFactor := 1
	i := 2

	for i*i <= n { //when i^2 > n: stop
		if n%i == 0 { //i is a factor of n
			largestFactor = i //i becomes the largest factor of n
			for n%i == 0 {    //get the lowest result of the division of n by factor i
				n /= i
			}
		}
		if i == 2 {
			i = 3
		} else { //add 2 to the counter and continue
			i += 2
		}
	}
	if n > 1 { //if n is bigger then 1, the largest factor is n
		largestFactor = n //else the largest factor is i
	}
	return largestFactor

}

func main() {
	//answer
	fmt.Println("Answer", maxPrimeFactor(600851475143))

	//below just having fun
	primeMap := make(map[int]int)

	for i := 2; i < 100000; i++ {
		if _, ok := primeMap[maxPrimeFactor(i)]; ok {
			primeMap[maxPrimeFactor(i)] = primeMap[maxPrimeFactor(i)] + 1
		} else {
			primeMap[maxPrimeFactor(i)] = 1
		}
	}

	var keys []int
	for k := range primeMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("highest prime factor %v, appears, %v times below 100000\n", k, primeMap[k])
	}

}
