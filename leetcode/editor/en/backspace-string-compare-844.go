package main

//Given two strings S and T, return if they are equal when both are typed into e
//mpty text editors. # means a backspace character.
//
// Note that after backspacing an empty text, the text will continue empty.
//
//
// Example 1:
//
//
//Input: S = "ab#c", T = "ad#c"
//Output: true
//Explanation: Both S and T become "ac".
//
//
//
// Example 2:
//
//
//Input: S = "ab##", T = "c#d#"
//Output: true
//Explanation: Both S and T become "".
//
//
//
// Example 3:
//
//
//Input: S = "a##c", T = "#a#c"
//Output: true
//Explanation: Both S and T become "c".
//
//
//
// Example 4:
//
//
//Input: S = "a#c", T = "b"
//Output: false
//Explanation: S becomes "c" while T becomes "b".
//
//
// Note:
//
//
// 1 <= S.length <= 200
// 1 <= T.length <= 200
// S and T only contain lowercase letters and '#' characters.
//
//
// Follow up:
//
//
// Can you solve it in O(N) time and O(1) space?
//
//
//
//
//
// Related Topics Two Pointers Stack
// 👍 2229 👎 105

//2021-02-05 17:01:09

//leetcode submit region begin(Prohibit modification and deletion)
func backspaceCompare_2(S string, T string) bool {
	i, j, back_s, back_t := len(S)-1, len(T)-1, 0, 0

	for true {
		for i >= 0 && (back_s > 0 || S[i] == '#') {
			if S[i] == '#' {
				back_s++
			} else {
				back_s--
			}
			i--
		}

		for j >= 0 && (back_t > 0 || T[j] == '#') {
			if T[j] == '#' {
				back_t++
			} else {
				back_t--
			}
			j--
		}

		if !(i >= 0 && j >= 0 && S[i] == T[j]) {
			return i == j && i == -1
		}
		i--
		j--

	}
	return false
}

func backspaceCompare(S string, T string) bool {
	s_ch, t_ch := make(chan bool), make(chan byte)
	go func(ch1 chan bool, ch2 chan byte) {
		back := 0
		for i := len(S) - 1; i >= 0; i-- {
			if S[i] == '#' {
				back++
			} else if back > 0 {
				back--
			} else {
				c, ok := <-ch2
				if !ok || (c != S[i]) {
					ch1 <- false
				}
			}
		}
		_, ok := <-ch2
		ch1 <- !ok
	}(s_ch, t_ch)

	go func(ch chan byte) {
		back := 0
		for i := len(T) - 1; i >= 0; i-- {
			if T[i] == '#' {
				back++
			} else if back > 0 {
				back--
			} else {
				ch <- T[i]
			}
		}
		close(ch)

	}(t_ch)

	return <-s_ch

}

//leetcode submit region end(Prohibit modification and deletion)
