package main

import "fmt"

func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // item at index 5 not included!

	// load the array with integers: 0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// print the slice
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// grow the slice beyond capacity
	// panic: runtime error: slice bound out of range
	//slice1 = slice1[0:7 ]

	// arr is a array, because [5] instead of []
	var arr = [5]int{0, 1, 2, 3, 4}
	// compile error, cannot use [5] int as [] int
	//sum_slice(arr)
	sum_slice(arr[:])

}

// a is a slice, []int, there is no length in '[]'
func sum_slice(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
