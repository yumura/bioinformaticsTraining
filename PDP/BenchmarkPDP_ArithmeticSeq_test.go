package pdp

import "testing"

func BenchmarkSPDP_ArithmeticSeq(b *testing.B) {
	helperPDP(SkienaPDP, makeASeq(), b)
}

func BenchmarkABFPDP_ArithmeticSeq(b *testing.B) {
	helperPDP(AnotherBruteForcePDP, makeASeq(), b)
}

func BenchmarkBFPDP_ArithmeticSeq(b *testing.B) {
	helperPDP(BruteForcePDP, makeASeq(), b)
}

func helperPDP(f func(*RMDelta, int) <-chan RestrictionMap, ints []int, b *testing.B) {
	answer := NewRestrictionMap(ints...).Normalize()
	delta := answer.Delta()
	n := answer.Length()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for result := range f(delta, n) {
			_ = result
		}
	}
}

func makeASeq() []int {
	n := 10
	d := 2

	xs := make([]int, n)
	for i := 0; i < n; i++ {
		xs[i] = i * d
	}
	return xs
}

// func TestSPDP_ArithmeticSeq(t *testing.T) {
// 	testBenchmarkPDP(SkienaPDP, makeASeq(), t)
// }
//
// func TestABFPDP_ArithmeticSeq(t *testing.T) {
// 	testBenchmarkPDP(AnotherBruteForcePDP, makeASeq(), t)
// }
//
// func TestBFPDP_ArithmeticSeq(t *testing.T) {
// 	testBenchmarkPDP(BruteForcePDP, makeASeq(), t)
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
