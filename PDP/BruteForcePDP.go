package pdp

import "sync"

// BruteForcePDP は 総当たりで PDP を解く
func BruteForcePDP(rmd *RMDelta, n int) <-chan RestrictionMap {
	ch, closed := makePDPChan(rmd, n)
	if closed {
		return ch
	}

	go bfPDP(rmd, n, ch)
	return ch
}

func bfPDP(rmd *RMDelta, n int, ch chan<- RestrictionMap) {
	m, _ := rmd.Max()
	width := int(m)

	wg := new(sync.WaitGroup)
	for xs := range CombinationInts(1, width-1, n-2) {
		ys := xs

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

func makePDPChan(rmd *RMDelta, n int) (ch chan RestrictionMap, closed bool) {
	ch = make(chan RestrictionMap)
	closed = true

	if rmd.Length() < 0 || rmd.Length() != n*(n-1)/2 {
		close(ch)
		return ch, closed
	}

	if rmd.Length() == 0 {
		go func() {
			ch <- *NewRestrictionMap()
			ch <- *NewRestrictionMap(0)
			close(ch)
		}()
		return ch, closed
	}

	if rmd.Length() == 1 {
		go func() {
			width := <-rmd.Chan()
			ch <- *NewRestrictionMap(0, int(width))
			close(ch)
		}()
		return ch, closed
	}

	closed = false
	return ch, closed
}
