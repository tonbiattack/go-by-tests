package sliceexample

import "slices"

func CloneThenChange(values []int) (original []int, cloned []int) {
	cloned = slices.Clone(values)
	cloned[0] = 99
	return values, cloned
}
