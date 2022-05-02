package main

import "fmt"

func confusingNumberII(n int) int {
	res := 0
	confusings := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}
	var dfs func(int, int, int)
	dfs = func(normalNum, confusingNum, decimal int) {
		if normalNum > n {
			return
		}
		if normalNum != confusingNum {
			fmt.Printf("%d ", normalNum)
			res++
		}
		for k, e := range confusings {
			if normalNum == 0 && k == 0 {
				continue
			}
			dfs(normalNum*decimal+k, e*decimal+confusingNum, decimal*10)
		}
		return

	}
	dfs(0, 0, 1)

	return res

}

func main() {
	res := confusingNumberII(20)
	println(res)
	res = confusingNumberII(100)
	println(res)
}
