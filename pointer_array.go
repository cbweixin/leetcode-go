package main

import "fmt"

func fp1(a [3]int)  { fmt.Println(a) }
func fp2(a *[3]int) { fmt.Println(a) }

func main() {
	a := [...]string{"a", "b", "c", "d"}
	b := []string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}
	for i := range b {
		fmt.Println("Array item", i, "is", b[i])
	}
	var ar [3]int
	fp1(ar)  // passes a copy of ar
	fp2(&ar) // passes a pointer to ar
}
