package main

import "fmt"

func main() {
	b()
}

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func aa() {
	trace("aa")
	defer untrace("aa")
	fmt.Println("in aa")
}

func b() {
	trace("b")
	defer untrace("b")
	fmt.Println("in b")
	aa()
}
