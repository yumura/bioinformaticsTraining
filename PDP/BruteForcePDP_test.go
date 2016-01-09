package pdp

import "testing"

func TestBFPDP_Empty(t *testing.T) {
	empty := NewRestrictionMap()
	one := NewRestrictionMap(0)
	delta := empty.Delta()

	i := 0
	for rMap := range BruteForcePDP(delta, 0) {
		if !rMap.Equal(empty) && !rMap.Equal(one) {
			t.Errorf("BruteForcePDP(%v, 0) should be {%v, %v}, but contains %v", delta, empty, one, rMap)
		}
		i++
	}
	if i == 0 {
		t.Errorf("BruteForcePDP(%v, 0) should be {%v, %v}, but empty set", delta, empty, one)
	}
}

func TestBFPDP_ButN(t *testing.T) {
	answer := NewRestrictionMap(1, 2, 3, 4).Normalize()
	delta := answer.Delta()

	for rMap := range BruteForcePDP(delta, 3) {
		t.Errorf("BruteForcePDP(%v, %v) should be empty, but contains %v", delta, 2, rMap)
	}

	for rMap := range BruteForcePDP(delta, 5) {
		t.Errorf("BruteForcePDP(%v, %v) should be empty, but contains %v", delta, 4, rMap)
	}
}

func TestBFPDP_onlyAnswer(t *testing.T) {
	answer1 := NewRestrictionMap(5, 10).Normalize()
	delta1 := answer1.Delta()
	n1 := answer1.Length()

	i := 0
	for rMap := range BruteForcePDP(delta1, n1) {
		if !rMap.Equal(answer1) {
			t.Errorf("BruteForcePDP(%v, %v) should be %v, but %v", delta1, n1, answer1, rMap)
		}
		i++
	}
	if i == 0 {
		t.Errorf("BruteForcePDP(%v, %v) should be {%v}, but empty set", delta1, n1, answer1)
	}

	answer2 := NewRestrictionMap(4, 6, 8).Normalize()
	delta2 := answer2.Delta()
	n2 := answer2.Length()

	j := 0
	for rMap := range BruteForcePDP(delta2, n2) {
		if !rMap.Equal(answer2) {
			t.Errorf("BruteForcePDP(%v, %v) should be {%v}, but %v", delta2, n2, answer2, rMap)
		}
		j++
	}
	if j == 0 {
		t.Errorf("BruteForcePDP(%v, %v) should be {%v}, but empty set", delta2, n2, answer2)
	}
}

func TestBFPDP_HomometricAnswer(t *testing.T) {
	answer1_1 := NewRestrictionMap(0, 1, 5)
	answer1_2 := answer1_1.Reflection().Normalize()
	delta1 := answer1_1.Delta() // == answer1_2.Delta()
	n1 := answer1_1.Length()

	i := 0
	for rMap := range BruteForcePDP(delta1, n1) {
		if !rMap.Equal(answer1_1) && !rMap.Equal(answer1_2) {
			t.Errorf("BruteForcePDP(%v, %v) should be {%v, %v}, but contains %v", delta1, n1, answer1_1, answer1_2, rMap)
		}
		i++
	}
	if i == 0 {
		t.Errorf("BruteForcePDP(%v, %v) should be {%v, %v}, but empty set", delta1, n1, answer1_1, answer1_2)
	}

	answer2_1 := NewRestrictionMap(0, 1, 3, 8, 9, 11, 12, 13, 15)
	answer2_2 := answer2_1.Reflection().Normalize()
	answer2_3 := NewRestrictionMap(0, 1, 3, 4, 5, 7, 12, 13, 15)
	answer2_4 := answer2_3.Reflection().Normalize()
	as := []*RestrictionMap{answer2_1, answer2_2, answer2_3, answer2_4}

	delta2 := answer2_1.Delta() // == answer2_x.Delta()
	n2 := answer2_1.Length()

	j := 0
	for rMap := range BruteForcePDP(delta2, n2) {
		flag := false
		for _, a := range as {
			flag = flag || rMap.Equal(a)
		}
		if !flag {
			t.Errorf("BruteForcePDP(%v, %v) should be {%v}, but contains %v", delta2, n2, as, rMap)
		}
		j++
	}
	if j == 0 {
		t.Errorf("BruteForcePDP(%v, %v) should be {%v}, but empty set", delta2, n2, as)
	}
}
