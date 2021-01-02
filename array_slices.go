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
	// two additional ways to create slices
	// make([]int, 50, 100)
	// new([100]int)[0:50]
	var slice2 []int = make([]int, 10)
	// load the array/slice:
	for i := 0; i < len(slice1); i++ {
		slice2[i] = 5 * i
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("\nThe length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	var slice3 []int = make([]int, 4)

	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 3
	slice3[3] = 4

	for ix, value := range slice3 {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
	fmt.Println(sum_slice(slice3))

	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, season := range seasons {
		fmt.Printf("Season %d is: %s\n", ix, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	slice4 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice4); i++ {
		slice4 = slice4[0 : i+1]
		slice4[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice4))
	}

	// print the slice:
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice4[i])
	}

	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

}

// a is a slice, []int, there is no length in '[]'
func sum_slice(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}
