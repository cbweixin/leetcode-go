package main

import "fmt"

// my first go code
func findSubstring(s string, words []string) []int {
	if len(words) == 0 {
		return []int{}
	}
	res := []int{}
	counter := map[string]int{}
	for _, w := range words {
		counter[w]++
	}

	length, totalLen := len(words[0]), len(words[0])*len(words)

	for i := 0; i < len(s)-totalLen+1; i++ {
		tmpCounter := map[string]int{}
		for j := 0; j < len(words); j++ {
			t := s[i+j*length : i+(j+1)*length]
			w, ok := counter[t]
			if !ok {
				break
			}
			tmpCounter[t]++
			if tmpCounter[t] > w {
				break
			}
			if j == len(words)-1 {
				res = append(res, i)
			}
		}
	}

	return res

}

func main() {
	s := "lingmindraboofooowingdingbarrwingmonkeypoundcake"
	words := []string{"fooo", "barr", "wing", "ding", "wing"}
	fmt.Println(findSubstring(s, words))

}
