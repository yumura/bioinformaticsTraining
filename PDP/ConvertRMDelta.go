package pdp

// DeltaFrom は rMap の各要素と整数 x の距離 RMDelta を生成する
func (rMap RestrictionMap) DeltaFrom(x int) *RMDelta {
	set := make(map[uint]uint)
	length := 0

	for y := range rMap.Chan() {
		if y > x {
			set[uint(y-x)]++
		} else {
			set[uint(x-y)]++
		}
		length++
	}

	return &RMDelta{
		set:    set,
		length: length,
	}
}

// Delta は RestrictionMap を RMDelta に変換する
func (rMap RestrictionMap) Delta() *RMDelta {
	rmd := NewRMDelta()

	for x := range rMap.Chan() {
		for k, v := range rMap.DeltaFrom(x).set {
			rmd.set[k] += v
			rmd.length += int(v)
		}
	}

	rmd.length = (rmd.length - int(rmd.Multiplicity(0))) / 2

	delete(rmd.set, 0)
	for k, v := range rmd.set {
		rmd.set[k] = v / 2
	}

	return rmd
}
