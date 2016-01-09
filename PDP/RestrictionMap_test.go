package pdp

import "testing"

func TestRestrictionMapChan(t *testing.T) {
	empty := NewRestrictionMap()
	for x := range empty.Chan() {
		t.Errorf("empty should be RestrictionMap{}, but RestrictionMap{%v}", x)
	}

	ints := []int{0, 2, 4, 7, 10}
	rMap := NewRestrictionMap(ints...)
	for x := range rMap.Chan() {
		contains := false
		for _, y := range ints {
			contains = contains || x == y
		}
		if !contains {
			t.Errorf("%v not in %v", x, ints)
		}
	}
}

func TestRestrictionMapEqual(t *testing.T) {
	empty := NewRestrictionMap()
	otherEmpty := NewRestrictionMap()
	if !empty.Equal(otherEmpty) || !otherEmpty.Equal(empty) {
		t.Errorf("NewRestrictionMap() should be %v", empty)
	}

	ints1 := []int{0, 1, 5}
	rMap1 := NewRestrictionMap(ints1...)
	other1 := NewRestrictionMap(ints1...)
	if !rMap1.Equal(other1) || !other1.Equal(rMap1) {
		t.Errorf("NewRestrictionMap(%v...) should be %v", ints1, rMap1)
	}

	reflectionInts := []int{0, -1, -5}
	rMap2 := NewRestrictionMap(reflectionInts...)
	if rMap1.Equal(rMap2) || rMap2.Equal(rMap1) {
		t.Errorf("%v should be %v, but NewRestrictionMap(%v...)", rMap1, rMap1, reflectionInts)
	}

	shiftInts := []int{3, 4, 8}
	rMap3 := NewRestrictionMap(shiftInts...)
	if rMap1.Equal(rMap3) || rMap3.Equal(rMap1) {
		t.Errorf("%v should be %v, but NewRestrictionMap(%v...)", rMap1, rMap1, shiftInts)
	}

	if rMap1.Equal(empty) || empty.Equal(rMap1) {
		t.Errorf("empty should be empty, but NewRestrictionMap(%v...)", ints1)
	}
}

func TestRestrictionMapShift(t *testing.T) {
	empty := NewRestrictionMap()
	if !empty.Shift(3).Equal(empty) {
		t.Errorf("empty.Shift(x) should be empty")
	}

	ints1 := []int{0, 1, 5}
	ints2 := []int{3, 4, 8}
	rMap1 := NewRestrictionMap(ints1...)
	rMap2 := NewRestrictionMap(ints2...)

	if !rMap1.Shift(3).Equal(rMap2) {
		t.Errorf("%v.Reflection(3) should be %v", rMap1, rMap2)
	}

	if !rMap1.Shift(3).Shift(-3).Equal(rMap1) {
		t.Errorf("%v.Shift(3).Shift(-3) should be %v", rMap1, rMap1)
	}
}

func TestRestrictionMapReflection(t *testing.T) {
	empty := NewRestrictionMap()
	if !empty.Reflection().Equal(empty) {
		t.Errorf("empty.Reflection() should be empty")
	}

	ints1 := []int{0, 1, 5}
	ints2 := []int{0, -1, -5}
	rMap1 := NewRestrictionMap(ints1...)
	rMap2 := NewRestrictionMap(ints2...)

	if !rMap1.Reflection().Equal(rMap2) {
		t.Errorf("%v.Reflection() should be %v", rMap1, rMap2)
	}

	if !rMap1.Reflection().Reflection().Equal(rMap1) {
		t.Errorf("%v.Reflection().Reflection() should be %v", rMap1, rMap1)
	}
}

func TestRestrictionMapIsNormalized(t *testing.T) {
	empty := NewRestrictionMap()
	normal := NewRestrictionMap(0, 2, 4, 7, 10)
	negative := normal.Shift(3)
	positive := normal.Shift(-3)

	if !empty.IsNormalized() {
		t.Errorf("%v.IsNormalized() should be true", empty)
	}

	if !normal.IsNormalized() {
		t.Errorf("%v.IsNormalized() should be true", normal)
	}

	if negative.IsNormalized() {
		t.Errorf("%v.IsNormalized() should be false", negative)
	}

	if positive.IsNormalized() {
		t.Errorf("%v.IsNormalized() should be false", positive)
	}
}

func TestRestrictionMapNormalize(t *testing.T) {
	empty := NewRestrictionMap()
	normal := NewRestrictionMap(0, 2, 4, 7, 10)
	negative := normal.Shift(3)
	positive := normal.Shift(-3)

	rMaps := []*RestrictionMap{empty, normal, negative, positive}
	for _, rMap := range rMaps {
		if !rMap.Normalize().IsNormalized() {
			t.Errorf("%v.Normalize() should be normalized", rMap)
		}
	}

	if !negative.Normalize().Equal(normal) {
		t.Errorf("%v should be %v", negative.Normalize(), normal)
	}

	if !positive.Normalize().Equal(normal) {
		t.Errorf("%v should be %v", positive.Normalize(), normal)
	}
}

func TestRestrictionMapLength(t *testing.T) {
	empty := NewRestrictionMap()
	if empty.Length() != 0 {
		t.Errorf("%v's length should be 0", empty)
	}

	rm := NewRestrictionMap(0, 2, 4, 7, 10, 10, 10)
	if rm.Length() != 5 {
		t.Errorf("%v's length should be 5", rm)
	}
}

func TestRestrictionMapAppend(t *testing.T) {
	empty := NewRestrictionMap()
	ten := NewRestrictionMap(10)
	rMap := NewRestrictionMap(-10, 10)

	ten2 := empty.Append(10)
	if !ten2.Equal(ten) {
		t.Errorf("%v.Append(10) should be %v, but %v", empty, ten, ten2)
	}

	ten3 := ten.Append(10)
	if !ten3.Equal(ten) {
		t.Errorf("%v.Append(10) should be %v, but %v", ten, ten, ten3)
	}

	rMap2 := ten.Append(-10)
	if !rMap2.Equal(rMap) {
		t.Errorf("%v.Append(10) should be %v, but %v", ten, rMap, rMap2)
	}
}

func TestRestrictionMapString(t *testing.T) {
	empty := NewRestrictionMap()
	eStr := empty.String()
	if eStr != "RestrictionMap{}" {
		t.Errorf("%v is not 'RestrictionMap{}'", eStr)
	}

	rm := NewRestrictionMap(0, 10)
	rmStr := rm.String()
	if rmStr != "RestrictionMap{0, 10}" && rmStr != "RestrictionMap{10, 0}" {
		t.Errorf("%v.String() should be 'RestrictionMap{0, 10}'", rmStr)
	}
}
