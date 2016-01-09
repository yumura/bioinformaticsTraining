package pdp

import (
	"math"
	"testing"
)

func BenchmarkSPDP_GeometricSeq(b *testing.B) {
	helperPDP(SkienaPDP, makeGSe(), b)
}

func BenchmarkABFPDP_GeometricSeq(b *testing.B) {
	helperPDP(AnotherBruteForcePDP, makeGSe(), b)
}

func BenchmarkBFPDP_GeometricSeq(b *testing.B) {
	helperPDP(BruteForcePDP, makeGSe(), b)
}

func makeGSe() []int {
	n := 5
	r := 2

	xs := make([]int, n)
	for i := 1; i < n; i++ {
		xs[i] = int(math.Pow(float64(r), float64(i)))
	}
	return xs
}

// func TestSPDP_GeometricSeq(t *testing.T) {
// 	testBenchmarkPDP(SkienaPDP, makeGSe(), t)
// }
//
// func TestABFPDP_GeometricSeq(t *testing.T) {
// 	testBenchmarkPDP(AnotherBruteForcePDP, makeGSe(), t)
// }
//
// func TestBFPDP_GeometricSeq(t *testing.T) {
// 	testBenchmarkPDP(BruteForcePDP, makeGSe(), t)
// }
//
// func testBenchmarkPDP(f func(*RMDelta, int) <-chan RestrictionMap, ints []int, t *testing.T) {
// 	answer := NewRestrictionMap(ints...).Normalize()
// 	delta := answer.Delta()
// 	n := answer.Length()
//
// 	i := 0
// 	success := false
// 	for result := range f(delta, n) {
// 		success = success || result.Equal(answer)
// 		i++
// 	}
// 	if !success {
// 		t.Errorf("Not Success; %v", i)
// 	}
// }
