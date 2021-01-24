package main

import "fmt"

func parseTernary(expression string) string {
	length := len(expression)
	level := 0
	for j := 1; j < length; j++ {
		if expression[j] == '?' {
			level++
		}
		if expression[j] == ':' {
			level--
		}
		if level == 0 {
			if expression[0] == 'T' {
				return parseTernary(expression[2:j])
			} else {
				return parseTernary(expression[j:length])
			}
		}
	}

	return expression
}

func main() {
	fmt.Println(parseTernary("T?T?F:5:3"))
	fmt.Println(parseTernary("F?1:T?4:5"))
}
