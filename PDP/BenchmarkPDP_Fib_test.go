package pdp

import "testing"

func BenchmarkSPDP_Fib(b *testing.B) {
	helperPDP(SkienaPDP, makeFibs(), b)
}

func BenchmarkABFPDP_Fib(b *testing.B) {
	helperPDP(AnotherBruteForcePDP, makeFibs(), b)
}

func BenchmarkBFPDP_Fib(b *testing.B) {
	helperPDP(BruteForcePDP, makeFibs(), b)
}

func makeFibs() []int {
	n := 8

	fibs := make([]int, n)
	fibs[0] = 0
	fibs[1] = 1
	for i := 2; i < n; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
	return fibs
}

// func TestSPDP_Fib(t *testing.T) {
// 	testBenchmarkPDP(SkienaPDP, makeFibs(), t)
// }
//
// func TestABFPDP_Fib(t *testing.T) {
// 	testBenchmarkPDP(AnotherBruteForcePDP, makeFibs(), t)
// }
//
// func TestBFPDP_Fib(t *testing.T) {
// 	testBenchmarkPDP(BruteForcePDP, makeFibs(), t)
// }
