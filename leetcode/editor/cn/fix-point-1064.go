package main

func fixedPoint1(arr []int) int {
	for i, v := range arr {
		if v == i {
			return i
		}
	}
	return -1
}

// binary_search
func fixedPoint(arr []int) int {
	var l int = 0
	var r int = len(arr) - 1

	for l < r {
		var mid int = (l + r) >> 1
		if arr[mid]-mid < 0 {
			l = mid + 1
		} else {
			r = mid
		}
	}

	if arr[l] == l {
		return l
	}
	return -1

}
