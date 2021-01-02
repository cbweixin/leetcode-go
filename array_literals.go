package main

import "fmt"

func main() {
	// var arrAge = [5]int{18, 20, 15, 22, 16}
	// var arrLazy = [...]int{5, 6, 7, 8, 22}
	// var arrLazy = []int{5, 6, 7, 8, 22}	//注：初始化得到的实际上是切片slice
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	// var arrKeyValue = []string{3: "Chris", 4: "Ron"}	//注：初始化得到的实际上是切片slice

	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}

	array := [3]float64{7.0, 8.5, 9.1}
	x := Sum_1(&array) // Note the explicit address-of operator
	// to pass a pointer to the array
	fmt.Printf("The sum of the array is: %f", x)
}

func Sum_1(a *[3]float64) (sum float64) {
	for _, v := range a { // derefencing *a to get back to the array is not necessary!
		sum += v
	}
	return
}
