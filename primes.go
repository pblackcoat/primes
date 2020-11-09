// Let's learn how to caluculate prime numbers
//Then, let's do that with multi-threading!

package primes

import (
	"fmt"
	"math"
)

func main() {
	var largestNumber int = 1000
	wa := BuildWorkArray(largestNumber)
	primes := Eratosthenes(wa)
	fmt.Printf("the following are all of the prime numbers %d or less: \n", largestNumber)
	for i := range primes {
		if i < len(primes)-1 {
			fmt.Printf("%d, ", primes[i])
		} else {
			fmt.Printf("%d\n", primes[i])
		}
	}
}

func Eratosthenes(wa []Number) []int {
	for i := range wa {
		if wa[i].prime == true {
			coefficient := 0
			for j := wa[i].value; j <= len(wa)+1; j++ {
				adder := wa[i].value * coefficient
				potential := int(math.Pow(float64(wa[i].value), 2)) + adder
				if j == potential {
					wa[j-2].prime = false
					coefficient++
				}
			}
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

func BuildWorkArray(highestValue int) []Number {
	var workArray []Number
	for i := 2; i <= highestValue; i++ {
		var val Number
		val.value = i
		val.prime = true
		workArray = append(workArray, val)
	}
	return workArray
}

func printSlice(s []int) {
	fmt.Printf("\nlen = %d cap=%d %v\n", len(s), cap(s), s)
}

type Number struct {
	value int
	prime bool
}
