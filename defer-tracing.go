package main

import "fmt"

func main() {
	b()
	bb()
}

func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

func trace1(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func ab() {
	defer un(trace1("ab"))
	fmt.Println("in ab")
}

func bb() {
	defer un(trace1("bb"))
	fmt.Println("in bb")
	ab()
}

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
