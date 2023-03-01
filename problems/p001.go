
// Solution to Project Euler problem 1
// Multiples of 3 or 5
//
// If we list all the natural numbers below 10 that are multiples
// of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
//
// Find the sum of all the multiples of 3 or 5 below 1000.
//

package main

import "fmt"

func main(){
	fmt.Println(sumOfMultiples3or5(1000))
}

func sumOfMultiples3or5(num_limit int) int {
	sum := 0
	for i := 0; i  <num_limit; i++ {
		if (i % 3 == 0 || i % 5 == 0){
			sum += i
		}
	}
	return sum
}