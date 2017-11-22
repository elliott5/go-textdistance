package textdistance

import "unsafe"

// Min returns the minimum number of passed int slices.
func Min(is ...int) int {
	var i int
	min := int(1<<(unsafe.Sizeof(i)*8-1) - 1)
	for _, v := range is {
		if min > v {
			min = v
		}
	}
	return min
}

// Max returns the maximum number of passed int slices.
func Max(is ...int) int {
	var max int
	for _, v := range is {
		if max < v {
			max = v
		}
	}
	return max
}
