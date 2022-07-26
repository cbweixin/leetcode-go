package main

import "fmt"

func badCall() {
	panic("runtime error: bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}

func fullName(firstName *string, lastName *string) {
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func main() {
	// fmt.Println("starting programm...")
	// panic("pacnic oh, oaa")
	// fmt.Println("hello")

	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")

	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")

}
