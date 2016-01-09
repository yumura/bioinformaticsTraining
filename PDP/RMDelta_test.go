package pdp

import "testing"

func TestRMDeltaChan(t *testing.T) {
	var ints = []uint{2, 2, 3, 3, 4, 5, 6, 7, 8, 10}

	var empty = NewRMDelta()
	var rmd1 = NewRMDelta(ints...)

	for x := range empty.Chan() {
		t.Errorf("empty should be RMDelta{}, but RMDelta{%v}", x)
	}

	for x := range rmd1.Chan() {
		contains := false
		for _, y := range ints {
			contains = contains || x == y
		}
		if !contains {
			t.Errorf("%v not in %v", x, ints)
		}
	}
}

func TestRMDeltaUniqueChan(t *testing.T) {
	var ints = []uint{2, 2, 3, 3, 4, 5, 6, 7, 8, 10}

	var empty = NewRMDelta()
	var rmd1 = NewRMDelta(ints...)

	for x := range empty.UniqueChan() {
		t.Errorf("empty should be RMDelta{}, but RMDelta{%v}", x)
	}

	flags := make(map[uint]bool)
	for x := range rmd1.UniqueChan() {
		if flags[x] {
			t.Errorf("%v not unique", x)
		} else {
			flags[x] = true
		}
	}
	for x := range flags {
		contains := false
		for _, y := range ints {
			contains = contains || x == y
		}
		if !contains {
			t.Errorf("%v not in %v", x, ints)
		}
	}
}

func TestRMDeltaEqual(t *testing.T) {
	var ints1 = []uint{2, 2, 3, 3, 4, 5, 6, 7, 8, 10}
	var ints2 = []uint{1, 1, 2, 1, 2, 3, 1, 2, 3, 4}

	var empty = NewRMDelta()
	var rmd1 = NewRMDelta(ints1...)
	var rmd2 = NewRMDelta(ints2...)

	otherEmpty := NewRMDelta()
	if !empty.Equal(otherEmpty) || !otherEmpty.Equal(empty) {
		t.Errorf("empty does not equal other empty")
	}

	otherRmd1 := NewRMDelta(ints1...)
	if !rmd1.Equal(otherRmd1) || !otherRmd1.Equal(rmd1) {
		t.Errorf("rmd1 does not equal other rmd1")
	}

	otherRmd2 := NewRMDelta(ints2...)
	if !rmd2.Equal(otherRmd2) || !otherRmd2.Equal(rmd2) {
		t.Errorf("rmd2 does not equal other rmd2")
	}

	rmds := []*RMDelta{empty, rmd1, rmd2}
	for i := range rmds {
		for j := range rmds {
			if i != j {
				if rmds[i].Equal(rmds[j]) {
					t.Errorf("%v equal %v", rmds[i], rmds[j])
				}
			}
		}
	}
}

func TestRMDeltaUnion(t *testing.T) {
	var ints1 = []uint{2, 2, 3, 3, 4, 5, 6, 7, 8, 10}
	var ints2 = []uint{1, 1, 2, 1, 2, 3, 1, 2, 3, 4}
	var ints3 = []uint{2, 2, 3, 3, 4, 5, 6, 7, 8, 10, 1, 1, 2, 1, 2, 3, 1, 2, 3, 4}

	var empty = NewRMDelta()
	var rmd1 = NewRMDelta(ints1...)
	var rmd2 = NewRMDelta(ints2...)
	var rmd3 = NewRMDelta(ints3...)

	if !empty.Union(rmd1).Equal(rmd1) || !rmd1.Union(empty).Equal(rmd1) {
		t.Errorf("empty should be identity element")
	}

	if rmd1.Union(rmd2).Equal(rmd1) {
		t.Errorf("Union() is not bang method")
	}

	if !rmd1.Union(rmd2).Equal(rmd3) || !rmd2.Union(rmd1).Equal(rmd3) {
		t.Errorf("%v union %v should be %v", rmd1, rmd2, rmd3)
	}
}

func TestRMDeltaDifference(t *testing.T) {
	var ints1 = []uint{2, 2, 3, 3, 4, 5, 6, 7, 8, 10}
	var ints2 = []uint{1, 1, 2, 1, 2, 3, 1, 2, 3, 4}
	var ints3 = []uint{5, 6, 7, 8, 10}
	var ints4 = []uint{1, 1, 1, 1, 2}

	var empty = NewRMDelta()
	var rmd1 = NewRMDelta(ints1...)
	var rmd2 = NewRMDelta(ints2...)
	var rmd3 = NewRMDelta(ints3...)
	var rmd4 = NewRMDelta(ints4...)

	if rmd1.Difference(rmd2).Equal(rmd1) {
		t.Errorf("Difference() is not bang method")
	}

	if !empty.Difference(rmd1).Equal(empty) {
		t.Errorf("empty difference other should be empty")
	}

	if !rmd1.Difference(empty).Equal(rmd1) {
		t.Errorf("other difference empty should be other")
	}

	if !rmd1.Difference(rmd2).Equal(rmd3) {
		t.Errorf("%v difference %v should be %v", rmd1, rmd2, rmd3)
	}

	if !rmd2.Difference(rmd1).Equal(rmd4) {
		t.Errorf("%v difference %v should be %v", rmd1, rmd2, rmd4)
	}
}

func TestRMDeltaLength(t *testing.T) {
	var ints = []uint{2, 2, 3, 3, 4, 5, 6, 7, 8, 10}

	var empty = NewRMDelta()
	var rmd1 = NewRMDelta(ints...)

	if empty.Length() != 0 {
		t.Errorf("empty's length should be zero")
	}

	if rmd1.Length() != len(ints) {
		t.Errorf("rmd1's length should be %v", len(ints))
	}
}

func TestRMDeltaMultiplicity(t *testing.T) {
	var ints = []uint{1, 1, 2, 1, 2, 3, 1, 2, 3, 4}

	var empty = NewRMDelta()
	var rmd1 = NewRMDelta(ints...)

	for k, v := range map[uint]uint{1: 4, 2: 3, 3: 2, 4: 1} {
		if empty.Multiplicity(k) != 0 {
			t.Errorf("empty.Multiplicity(%v) should be zero", k)
		}

		if rmd1.Multiplicity(k) != v {
			t.Errorf("rmd1.Multiplicity(%v) should be %v", k, v)
		}
	}
}

func TestRMDeltaMax(t *testing.T) {
	var ints1 = []uint{2, 2, 3, 3, 4, 5, 6, 7, 8, 10}
	var ints2 = []uint{1, 1, 2, 1, 2, 3, 1, 2, 3, 4}

	var empty = NewRMDelta()
	var rmd1 = NewRMDelta(ints1...)
	var rmd2 = NewRMDelta(ints2...)

	max, ok := empty.Max()
	if max != 0 || ok {
		t.Errorf("empty.max() should be (0, NG)")
	}

	max, ok = rmd1.Max()
	if max != 10 || !ok {
		t.Errorf("rmd1.max() should be (10, OK)")
	}

	max, ok = rmd2.Max()
	if max != 4 || !ok {
		t.Errorf("rmd2.max() should be (4, OK)")
	}
}

func TestRMDeltaString(t *testing.T) {
	empty := NewRMDelta()
	eStr := empty.String()
	if eStr != "RMDelta{}" {
		t.Errorf("%v is not 'RMDelta{}'", eStr)
	}

	rm := NewRMDelta(0, 10, 10)
	rmStr := rm.String()
	if rmStr != "RMDelta{0, 10, 10}" && rmStr != "RMDelta{10, 0, 10}" && rmStr != "RMDelta{10, 10, 0}" {
		t.Errorf("%v.String() should be 'RMDelta{0, 10, 10}'", rmStr)
	}
}
