package pdp

import (
	"reflect"
	"testing"
)

func TestCombinationInts_BatArgs(t *testing.T) {
	_, ok1 := <-CombinationInts(0, 10, -5)
	if ok1 {
		t.Errorf("CombinationInts(0, 10, -10) should be closed channel, but open")
	}

	_, ok2 := <-CombinationInts(0, -10, 5)
	if ok2 {
		t.Errorf("CombinationInts(0, -10, 5) should be closed channel, but open")
	}

	_, ok3 := <-CombinationInts(0, 5, 10)
	if ok3 {
		t.Errorf("CombinationInts(0, 5, 10) should be closed channel, but open")
	}
}

func TestCombinationInts_Normal(t *testing.T) {
	answer := [][]int{[]int{0, 1, 2}, []int{0, 1, 3}, []int{0, 2, 3}, []int{1, 2, 3}}
	result := make([][]int, 4)

	i := 0
	for ints := range CombinationInts(0, 4, 3) {
		result[i] = ints
		i++
	}

	if !reflect.DeepEqual(result, answer) {
		t.Errorf("CombinationInts(0, 4, 3) should be %v, but %v", answer, result)
	}
}

func TestCombinationInts_Shift(t *testing.T) {
	answer := [][]int{[]int{10, 11, 12}, []int{10, 11, 13}, []int{10, 12, 13}, []int{11, 12, 13}}
	result := make([][]int, 4)

	i := 0
	for ints := range CombinationInts(10, 4, 3) {
		result[i] = ints
		i++
	}

	if !reflect.DeepEqual(result, answer) {
		t.Errorf("CombinationInts(10, 4, 3) should be %v, but %v", answer, result)
	}
}
