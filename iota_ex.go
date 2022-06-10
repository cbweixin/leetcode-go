package main

import "fmt"

const (
	Sunday    = iota // starting from 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1
	MB                          // MB = 1 << (10*iota) but iota = 2
	GB                          // same as GB, but iota is 3
	TB                          // iota = 4
)

const (
	AA, BB = iota, iota << 10 // 0, 0 << 10
	CC, DD                    // 1, 1 << 10
)

const (
	A1 = iota // 0
	B1        // 1
	C  = "c"  // c
	D         // c
	E  = iota // 4
	F         // 5
)

func main() {
	fmt.Println(Sunday)
	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Wednesday)
	fmt.Println(Thursday)
	fmt.Println(Friday)
	fmt.Println(Saturday)

	println(KB)
	println(MB)
	println(GB)
	println(TB)

	println(AA)
	println(BB)
	println(CC)
	println(DD)
}
