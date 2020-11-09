// Let's learn how to caluculate prime numbers
//Then, let's do that with multi-threading!

package main

import (
	"fmt"
	"math"
)

func main() {
	var largestNumber int = 100

	wa := buildWorkArray(largestNumber)
	for i := range wa {
		if wa[i].prime == true {
			coefficient := 0
			for j := wa[i].value; j <= largestNumber; j++ {
				adder := wa[i].value * coefficient
				fmt.Printf("i = %d, j = %d, adder = %d \n", wa[i].value, j, adder)
				potential := int(math.Pow(float64(wa[i].value), 2)) + adder
				if j == potential {
					fmt.Printf("Match, %d is a multiple of %d\n", potential, wa[i].value)
					wa[j-2].prime = false
					coefficient++
				}
			}
		}
	}
	printSlice(wa)
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

func printSlice(s []number) {
	fmt.Printf("\nlen = %d cap=%d %v\n", len(s), cap(s), s)
}

type number struct {
	value int
	prime bool
}
