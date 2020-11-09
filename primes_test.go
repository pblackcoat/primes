package main

import (
	"reflect"
	"testing"
)

func TestEratosthenes(t *testing.T) {
	got := eratosthenes(buildWorkArray(10))
	testArray := []int{2, 3, 5, 7}
	var ts []int = testArray[0:len(testArray)]
	if reflect.DeepEqual(got, ts) {

	} else {
		t.Fail()
	}
}

func BenchmarkBuildWorkArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buildWorkArray(10000)
	}
}

func BenchmarkSieve(b *testing.B) {
	big := buildWorkArray(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sieve(big, 2)
	}
}
