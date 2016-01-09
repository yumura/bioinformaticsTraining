package pdp

import "testing"

func TestDeltaFrom(t *testing.T) {
	helperDeltaFrom(NewRestrictionMap(), 5, NewRMDelta(), t)

	rMap := NewRestrictionMap(0, 2, 4)
	rmd1 := NewRMDelta(1, 1, 3)
	helperDeltaFrom(rMap, 1, rmd1, t)
	helperDeltaFrom(rMap.Reflection(), -1, rmd1, t)
	helperDeltaFrom(rMap.Shift(5), 1+5, rmd1, t)
}

func helperDeltaFrom(rMap *RestrictionMap, x int, rmd *RMDelta, t *testing.T) {
	rmd2 := rMap.DeltaFrom(x)
	if !rmd2.Equal(rmd) {
		t.Errorf("%v.DeltaFrom(%v) should be %v, but %v", rMap, x, rmd, rmd2)
	}
}

func TestDelta(t *testing.T) {
	helperDelta(NewRestrictionMap(), NewRMDelta(), t)
	helperDelta(NewRestrictionMap(1), NewRMDelta(), t)

	rMap := NewRestrictionMap(0, 2, 4, 7, 10)
	rmd1 := NewRMDelta(2, 2, 3, 3, 4, 5, 6, 7, 8, 10)
	helperDelta(rMap, rmd1, t)
	helperDelta(rMap.Reflection(), rmd1, t)
	helperDelta(rMap.Shift(5), rmd1, t)
	helperDelta(rMap.Shift(-5), rmd1, t)
}

func helperDelta(rMap *RestrictionMap, rmd *RMDelta, t *testing.T) {
	rmd2 := rMap.Delta()
	if !rmd2.Equal(rmd) {
		t.Errorf("%v.Delta() should be (%v, %v), but (%v, %v)", rMap, rmd, rmd.Length(), rmd2, rmd2.Length())
	}
}
