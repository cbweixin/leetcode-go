package main

import (
	"sort"
	"strconv"
)

//
//A self-dividing number is a number that is divisible by every digit it contain
//s.
//
//For example, 128 is a self-dividing number because 128 % 1 == 0, 128 % 2 == 0,
// and 128 % 8 == 0.
//
//Also, a self-dividing number is not allowed to contain the digit zero.
//
//Given a lower and upper number bound, output a list of every possible self div
//iding number, including the bounds if possible.
//
// Example 1:
//
//Input:
//left = 1, right = 22
//Output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22]
//
//
//
// Note:
// The boundaries of each input argument are 1 <= left <= right <= 10000.
// Related Topics Math
// 👍 785 👎 304

// 2021-01-11 21:54:38

//leetcode submit region begin(Prohibit modification and deletion)
func selfDividingNumbers(left int, right int) []int {
	type Response struct {
		Num   int
		Valid bool
	}

	// a buffered channel, so no block
	ch := make(chan Response, right-left+1)

	// For every input, spin up a goroutine to evaluate whether the number is a solution.
	// Stick a response in the channel.
	for i := left; i <= right; i++ {
		go func(ch chan Response, num int) {
			numString := strconv.Itoa(num)
			for j := range numString {
				digit, _ := strconv.Atoi(string(numString[j]))
				if digit == 0 || num%digit != 0 {
					ch <- Response{num, false}
					return
				}
			}
			ch <- Response{num, true}

		}(ch, i)
	}

	res := make([]int, 0)
	for i := left; i <= right; i++ {
		resp := <-ch
		if resp.Valid {
			res = append(res, resp.Num)
		}
	}

	sort.Ints(res)

	return res

}

//leetcode submit region end(Prohibit modification and deletion)
