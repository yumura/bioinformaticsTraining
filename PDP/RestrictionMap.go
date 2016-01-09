package pdp

import (
	"strconv"
	"strings"
)

// RestrictionMap は制限酵素地図を表すための構造体
type RestrictionMap struct {
	set map[int]bool
	min int
}

// NewRestrictionMap は RestrictionMap のコンストラクタ
func NewRestrictionMap(ints ...int) *RestrictionMap {
	rMap := RestrictionMap{
		set: make(map[int]bool),
		min: 0,
	}

	for _, x := range ints {
		if rMap.min > x || len(rMap.set) == 0 {
			rMap.min = x
		}

		rMap.set[x] = true
	}

	return &rMap
}

// Chan は要素を返すチャンネルを返す
func (rMap RestrictionMap) Chan() <-chan int {
	ch := make(chan int)
	go func() {
		for k := range rMap.set {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

// Equal は RMDeltaの同値性を検査する
func (rMap RestrictionMap) Equal(rMap2 *RestrictionMap) bool {
	if rMap.min != rMap2.min || len(rMap.set) != len(rMap2.set) {
		return false
	}

	for x := range rMap.set {
		_, ok := rMap2.set[x]
		if !ok {
			return false
		}
	}

	return true
}

// Shift は各要素を x だけずらす
func (rMap RestrictionMap) Shift(x int) *RestrictionMap {
	if x == 0 || len(rMap.set) == 0 {
		return &RestrictionMap{
			set: rMap.set,
			min: rMap.min,
		}
	}

	set := make(map[int]bool)
	for y := range rMap.Chan() {
		set[y+x] = true
	}

	return &RestrictionMap{
		set: set,
		min: rMap.min + x,
	}
}

// Reflection は各要素の符号を反転する
func (rMap RestrictionMap) Reflection() *RestrictionMap {
	set := make(map[int]bool)
	min := -rMap.min
	for y := range rMap.Chan() {
		set[-y] = true
		if -y < min {
			min = -y
		}
	}

	return &RestrictionMap{
		set: set,
		min: min,
	}
}

// IsNormalized は正規化されているか調べる
func (rMap RestrictionMap) IsNormalized() bool {
	return len(rMap.set) == 0 || rMap.min == 0
}

// Normalize すると最小値が 0 になる
func (rMap RestrictionMap) Normalize() *RestrictionMap {
	return rMap.Shift(-rMap.min)
}

// Length は要素数を返す
func (rMap RestrictionMap) Length() int {
	return len(rMap.set)
}

// Append は要素を追加した RestrictionMap を返す
func (rMap RestrictionMap) Append(x int) *RestrictionMap {
	if rMap.set[x] {
		return &RestrictionMap{
			set: rMap.set,
			min: rMap.min,
		}
	}

	set := make(map[int]bool)
	set[x] = true
	for y := range rMap.Chan() {
		set[y] = true
	}

	min := x
	if rMap.Length() != 0 && rMap.min < min {
		min = rMap.min
	}

	return &RestrictionMap{
		set: set,
		min: min,
	}
}

func (rMap RestrictionMap) String() string {
	l := make([]string, 0, len(rMap.set))
	for x := range rMap.Chan() {
		l = append(l, strconv.Itoa(x))
	}
	return "RestrictionMap{" + strings.Join(l, ", ") + "}"
}
