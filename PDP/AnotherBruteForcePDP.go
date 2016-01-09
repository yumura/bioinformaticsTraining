package pdp

import "sync"

// AnotherBruteForcePDP は
// rmd 内の要素を用い
// 総当たりで PDP を解く
func AnotherBruteForcePDP(rmd *RMDelta, n int) <-chan RestrictionMap {
	ch, closed := makePDPChan(rmd, n)
	if closed {
		return ch
	}

	go abfPDP(rmd, n, ch)
	return ch
}

func abfPDP(rmd *RMDelta, n int, ch chan<- RestrictionMap) {
	m, _ := rmd.Max()
	width := int(m)
	us := uniqueInts(rmd)

	wg := new(sync.WaitGroup)
	for xs := range CombinationInts(0, len(us), n-2) {
		ys := mapping(xs, us)
		wg.Add(1)

		go func() {
			rMap := NewRestrictionMap(append(ys, 0, width)...)
			if rMap.Delta().Equal(rmd) {
				ch <- *rMap
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(ch)
}

func uniqueInts(rmd *RMDelta) []int {
	u := rmd.UniqueLength()

	us := make([]int, 0, u)
	for x := range rmd.UniqueChan() {
		us = append(us, int(x))
	}

	return us
}

func mapping(xs []int, us []int) []int {
	ys := make([]int, len(xs))
	for i, x := range xs {
		ys[i] = us[x]
	}
	return ys
}
