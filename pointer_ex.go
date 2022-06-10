package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := 0x12345678
	p := unsafe.Pointer(&x)
	n := (*[4]byte)(p)

	for i := 0; i < len(n); i++ {
		fmt.Printf("%X\n", n[i])
	}

	d := struct {
		s string
		x int
	}{"abc", 100}

	p1 := uintptr(unsafe.Pointer(&d))
	p1 += unsafe.Offsetof(d.x)
	p2 := unsafe.Pointer(p1)
	px := (*int)(p2)
	*px = 200
	fmt.Printf("%#v\n", d)
}
