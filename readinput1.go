package main

import "fmt"

var (
	firstName, lastName, s string
	ii                     int
	f_f                    float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func main() {

	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&firstName, &lastName)
	// fmt.Scanf("%s %s", &firstName, &lastName)
	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	fmt.Sscanf(input, format, &f_f, &ii, &s)
	fmt.Println("From the string we read: ", f_f, ii, s)
	// 输出结果: From the string we read: 56.12 5212 Go

}
