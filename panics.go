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
	// below line could be output because in test function, there is a defer which recovered the panic
	fmt.Printf("Test completed\r\n")

	firstName := "Elon"
	fullName(&firstName, nil)
	// 当出现了 panic 时，程序就会终止运行，打印出传入 panic 的参数，接着打印出堆栈跟踪
	// so below line won't be output

	fmt.Println("returned normally from main")

	// 我们重新总结一下 panic 做了什么。当函数发生 panic 时，它会终止运行，在执行完所有的延迟函数后，程序控制返回到该函数的调用方。这样的过
	// 程会一直持续下去，直到当前协程的所有函数都返回退出，然后程序会打印出 panic 信息，接着打印出堆栈跟踪，最后程序终止。

}
