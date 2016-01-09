package pdp

import (
	"strconv"
	"strings"
)

// RMDelta は制限酵素地図の各点の差（距離）の多重集合を表す
type RMDelta struct {
	set    map[uint]uint
	length int
}

// NewRMDelta は RestrictionMap のコンストラクタ
func NewRMDelta(ints ...uint) *RMDelta {
	rmd := RMDelta{
		set:    make(map[uint]uint),
		length: 0,
	}

	for _, x := range ints {
		rmd.set[x]++
		rmd.length++
	}

	return &rmd
}

// Chan は要素を返すチャンネルを返す
func (rmd RMDelta) Chan() <-chan uint {
	ch := make(chan uint)
	go func() {
		for k, v := range rmd.set {
			for i := uint(0); i < v; i++ {
				ch <- k
			}
		}
		close(ch)
	}()
	return ch
}

// UniqueChan は重複無く要素を返すチャンネルを返す
func (rmd RMDelta) UniqueChan() <-chan uint {
	ch := make(chan uint)
	go func() {
		for k := range rmd.set {
			ch <- k
		}
		close(ch)
	}()
	return ch
}

// Equal は RMDeltaの同値性を検査する
func (rmd RMDelta) Equal(rmd2 *RMDelta) bool {
	if rmd.length != rmd2.length || len(rmd2.set) != len(rmd2.set) {
		return false
	}

	for k, v := range rmd.set {
		v2, ok := rmd2.set[k]
		if !ok || v != v2 {
			return false
		}
	}

	return true
}

// Union は集合和
func (rmd RMDelta) Union(rmd2 *RMDelta) *RMDelta {
	newRmd := rmd.copy()

	for k, v := range rmd2.set {
		newRmd.set[k] += v
	}

	newRmd.length += rmd2.length

	return newRmd
}

// Difference は集合差
func (rmd RMDelta) Difference(rmd2 *RMDelta) *RMDelta {
	newRmd := rmd.copy()

	i := uint(0)

	for k, v := range rmd2.set {
		if newRmd.set[k] <= v {
			i += newRmd.set[k]
			delete(newRmd.set, k)
		} else {
			i += v
			newRmd.set[k] -= v
		}
	}

	newRmd.length -= int(i)

	return newRmd
}

func (rmd RMDelta) copy() *RMDelta {
	newRmd := NewRMDelta()
	for k, v := range rmd.set {
		newRmd.set[k] = v
	}
	newRmd.length = rmd.length
	return newRmd
}

// Length は要素数を返す
func (rmd RMDelta) Length() int {
	return rmd.length
}

// UniqueLength は重複をなくした場合の要素数を返す
func (rmd RMDelta) UniqueLength() int {
	return len(rmd.set)
}

// Multiplicity は要素 x の重複度を返す
// ただし、含まれていなければ 0 を返す
func (rmd RMDelta) Multiplicity(x uint) uint {
	return rmd.set[x]
}

// Max は要素の中の最大値を返す
func (rmd RMDelta) Max() (max uint, ok bool) {
	for k := range rmd.set {
		ok = true

		if k > max {
			max = k
		}
	}

	return max, ok
}

func (rmd RMDelta) String() string {
	l := make([]string, 0, len(rmd.set))
	for x := range rmd.Chan() {
		l = append(l, strconv.Itoa(int(x)))
	}
	return "RMDelta{" + strings.Join(l, ", ") + "}"
}
