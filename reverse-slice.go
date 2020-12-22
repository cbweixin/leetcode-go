package main

import (
	"fmt"
	"reflect"
)

//Reverse reverses a slice.
func Reverse(slice interface{}) {
	size := reflect.ValueOf(slice).Len()
	swap := reflect.Swapper(slice)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func main() {
	println("Please edit main.go,and complete the 'Reverse' function to pass the test.\nYou should use reflect " +
		"package to reflect the slice type and make it applly to any type.\nTo run test,please run 'go test'\nIf you" +
		" pass the test,please run 'git checkout l2' ")
	a := []int{48, 44, 37, 40, 14, 45, 8, 38, 12, 3}
	Reverse(a)
	fmt.Println(a)

}
