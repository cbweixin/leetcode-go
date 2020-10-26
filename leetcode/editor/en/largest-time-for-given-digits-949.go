package main

import "fmt"

//Given an array arr of 4 digits, find the latest 24-hour time that can be made
//using each digit exactly once.
//
// 24-hour times are formatted as "HH:MM", where HH is between 00 and 23, and MM
// is between 00 and 59. The earliest 24-hour time is 00:00, and the latest is 23:
//59.
//
// Return the latest 24-hour time in "HH:MM" format. If no valid time can be mad
//e, return an empty string.
//
//
// Example 1:
//
//
//Input: A = [1,2,3,4]
//Output: "23:41"
//Explanation: The valid 24-hour times are "12:34", "12:43", "13:24", "13:42", "
//14:23", "14:32", "21:34", "21:43", "23:14", and "23:41". Of these times, "23:41"
// is the latest.
//
//
// Example 2:
//
//
//Input: A = [5,5,5,5]
//Output: ""
//Explanation: There are no valid 24-hour times as "55:55" is not valid.
//
//
// Example 3:
//
//
//Input: A = [0,0,0,0]
//Output: "00:00"
//
//
// Example 4:
//
//
//Input: A = [0,0,1,0]
//Output: "10:00"
//
//
//
// Constraints:
//
//
// arr.length == 4
// 0 <= arr[i] <= 9
//
// Related Topics Math
// 👍 407 👎 789

// 2020-10-25 14:12:28

//leetcode submit region begin(Prohibit modification and deletion)
func largestTimeFromDigits(arr []int) string {
	ans := -1
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i != j {
				for k := 0; k < 4; k++ {
					if k != i && k != j {
						v := 6 - i - j - k
						hour, minute := 10*arr[i]+arr[j], 10*arr[k]+arr[v]
						if hour < 24 && minute < 60 {
							if hour*60+minute > ans {
								ans = hour*60 + minute
							}
						}
					}
				}
			}

		}
	}

	if ans >= 0 {
		return fmt.Sprintf("%02d:%02d", ans/60, ans%60)
	}

	return ""

}

//leetcode submit region end(Prohibit modification and deletion)
