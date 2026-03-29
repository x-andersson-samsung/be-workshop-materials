package adder

import (
	"math"
	"strconv"
	"testing"
)

// For future:
// You can use benchstat tool to compare benchmark results
// go install golang.org/x/perf/cmd/...@latest

//go test -run '^$' -bench '^BenchmarkAdd$' -benchmem ./... > old.txt
//go test -run '^$' -bench '^BenchmarkAdd2$' -benchmem ./... > new.txt
//
//Benchstat requires benchmark names to be the same, so we need to normalize them in Add2 file:
//sed 's/BenchmarkAdd2/BenchmarkAdd/g' new.txt > new-normalized.txt
//
//benchstat old.txt new-normalized.txt

func BenchmarkAdd(b *testing.B) {
	for count := 0; count < 8; count++ {
		number := int(math.Pow(10, float64(count)))
		b.Run(strconv.Itoa(number), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Add(number, number)
			}
		})
	}
}
