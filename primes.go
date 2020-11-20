// Let's learn how to caluculate prime numbers
//Then, let's do that with multi-threading!

package main

import (
	"fmt"
	"math"
)

func main() {
	var largestNumber int = 726000
	wa := buildWorkArray(largestNumber)
	primes := eratosthenesGr(wa)
	fmt.Printf("the following are all of the prime numbers %d or less: \n", largestNumber)
	for i := range primes {
		if i < len(primes)-1 {
			fmt.Printf("%d, ", primes[i])
		} else {
			fmt.Printf("%d\n", primes[i])
		}
	}
}

func eratosthenes(wa []number) []int {
	for i := range wa {
		if wa[i].prime == true {
			sieve(wa, wa[i].value)
		}
	}
	var primes []int
	for i := range wa {
		if wa[i].prime == true {
			primes = append(primes, wa[i].value)
		}
	}
	return primes
}

func eratosthenesGr(wa []number) []int {
	for i := range wa {
		if wa[i].prime == true {
			go sieve(wa, wa[i].value)
		}
	}
	var primes []int
	for i := range wa {
		if wa[i].prime == true {
			primes = append(primes, wa[i].value)
		}
	}
	return primes
}

func sieve(wa []number, filter int) []number {
	coefficient := 0
	for j := filter; j <= len(wa)+1; j++ {
		adder := filter * coefficient
		potential := int(math.Pow(float64(filter), 2)) + adder
		if j == potential {
			wa[j-2].prime = false
			coefficient++
		}
	}
	return wa
}

func buildWorkArray(highestValue int) []number {
	var workArray []number
	for i := 2; i <= highestValue; i++ {
		var val number
		val.value = i
		val.prime = true
		workArray = append(workArray, val)
	}
	return workArray
}

func printSlice(s []int) {
	fmt.Printf("\nlen = %d cap=%d %v\n", len(s), cap(s), s)
}

type number struct {
	value int
	prime bool
}
