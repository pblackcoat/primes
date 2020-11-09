package primes

import (
	"reflect"
	"testing"
)

func TestEratosthenes(t *testing.T) {
	got := Eratosthenes(BuildWorkArray(10))
	testArray := []int{2, 3, 5, 7}
	var ts []int = testArray[0:len(testArray)]
	if reflect.DeepEqual(got, ts) {

	} else {
		t.Fail()
	}
}
