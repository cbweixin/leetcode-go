package main

import "fmt"

func main() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}

	t := "abc"
	fmt.Println(t)
	t = t[:1] + "d" + t[2:]
	fmt.Println(t)

}
