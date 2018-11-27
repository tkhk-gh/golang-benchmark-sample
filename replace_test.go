package main

import (
	"testing"
)

func BenchmarkReplaceWithStrings(b *testing.B) {
	s := "test1.test2.test3.test4"
	for i := 0; i < b.N; i++ {
		replaceWithStrings(s)
	}
}

func BenchmarkReplaceWithRegex(b *testing.B) {
	s := "test1.test2.test3.test4"
	for i := 0; i < b.N; i++ {
		replaceWithRegex(s)
	}
}
