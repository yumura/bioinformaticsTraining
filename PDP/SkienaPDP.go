package pdp

import "sync"

// SkienaPDP は
// バックトラックを用い PDP を解く
func SkienaPDP(rmd *RMDelta, n int) <-chan RestrictionMap {
	ch, closed := makePDPChan(rmd, n)
	if closed {
		return ch
	}

	go sPDP(rmd, n, ch)
	return ch
}

func sPDP(rmd *RMDelta, n int, ch chan<- RestrictionMap) {
	width, _ := rmd.Max()

	if rmd.Multiplicity(width) != 1 {
		close(ch)
		return
	}

	newRmd := rmd.Difference(NewRMDelta(width))
	rMap := NewRestrictionMap(0, int(width))

	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		sPlace(newRmd, rMap, width, ch, wg)
		wg.Done()
	}()
	wg.Wait()
	close(ch)
}

func sPlace(rmd *RMDelta, rMap *RestrictionMap, width uint, ch chan<- RestrictionMap, wg *sync.WaitGroup) {
	n := rmd.Length()
	m, _ := rmd.Max()
	mm := rmd.Multiplicity(m)

	if n == 0 {
		ch <- *rMap
		return
	}

	if mm > 2 {
		return
	}

	if mm < 2 {
		wg.Add(1)
		go func() {
			rmdR, rMapR, okR := nextRM(rmd, rMap, m)
			if okR {
				sPlace(rmdR, rMapR, width, ch, wg)
			}
			wg.Done()
		}()

		wg.Add(1)
		go func() {
			rmdL, rMapL, okL := nextRM(rmd, rMap, width-m)
			if okL {
				sPlace(rmdL, rMapL, width, ch, wg)
			}
			wg.Done()
		}()

		return
	}

	rmdR, rMapR, okR := nextRM(rmd, rMap, m)
	if !okR {
		return
	}

	if m == width-m {
		wg.Add(1)
		go func() {
			sPlace(rmdR, rMapR, width, ch, wg)
			wg.Done()
		}()
		return
	}

	rmdL, rMapL, okL := nextRM(rmdR, rMapR, width-m)
	if okL {
		wg.Add(1)
		go func() {
			sPlace(rmdL, rMapL, width, ch, wg)
			wg.Done()
		}()
	}
	return
}

func nextRM(rmd *RMDelta, rMap *RestrictionMap, x uint) (*RMDelta, *RestrictionMap, bool) {
	df := rMap.DeltaFrom(int(x))
	nextRmd := rmd.Difference(df)

	if rmd.Length()-df.Length() == nextRmd.Length() {
		return nextRmd, rMap.Append(int(x)), true
	}

	return nil, nil, false
}
