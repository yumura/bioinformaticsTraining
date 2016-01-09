package pdp

// CombinationInts は
// m から連続する n 個の整数を
// 重複無く k 個選んだ時の組み合わせを列挙する
func CombinationInts(m int, n int, k int) <-chan []int {
	ch := make(chan []int)

	if 0 <= k && k <= n {
		go combInts(m, n, k, ch)
	} else {
		close(ch)
	}

	return ch
}

func combInts(m int, n int, k int, ch chan<- []int) {
	defer close(ch)

	xs := makeFirstConb(m, n, k)
	sendCopy(xs, ch)

	for xs[0] < m+n-k {
		updateNextConb(xs, k)
		sendCopy(xs, ch)
	}
}

func makeFirstConb(m int, n int, k int) []int {
	xs := make([]int, k+1)
	xs[0] = m
	for i := 1; i < k; i++ {
		xs[i] = xs[i-1] + 1
	}
	xs[k] = m + n
	return xs
}

func updateNextConb(xs []int, k int) {
	for i := k - 1; i >= 0; i-- {
		if xs[i]+1 < xs[i+1] {
			xs[i]++

			for j := i + 1; j < k; j++ {
				xs[j] = xs[j-1] + 1
			}

			return
		}
	}
}

func sendCopy(xs []int, ch chan<- []int) {
	ys := make([]int, cap(xs)-1)
	copy(ys, xs[:len(xs)-1])
	ch <- ys
}
