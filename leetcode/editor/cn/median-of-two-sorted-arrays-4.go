package main

//给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
//
// 算法的时间复杂度应该为 O(log (m+n)) 。
//
//
//
// 示例 1：
//
//
//输入：nums1 = [1,3], nums2 = [2]
//输出：2.00000
//解释：合并数组 = [1,2,3] ，中位数 2
//
//
// 示例 2：
//
//
//输入：nums1 = [1,2], nums2 = [3,4]
//输出：2.50000
//解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
//
//
//
//
//
//
// 提示：
//
//
// nums1.length == m
// nums2.length == n
// 0 <= m <= 1000
// 0 <= n <= 1000
// 1 <= m + n <= 2000
// -10⁶ <= nums1[i], nums2[i] <= 10⁶
//
//
// Related Topics 数组 二分查找 分治 👍 6213 👎 0

//2023-01-27 14:09:20

// leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	var getKth func([]int, int, int, []int, int, int, int) int
	min := func(x, y int) int {
		if x <= y {
			return x
		}
		return y
	}

	// partition
	getKth = func(arr1 []int, start1 int, end1 int, arr2 []int, start2 int, end2 int, k int) int {
		l1, l2 := end1-start1+1, end2-start2+2
		if l1 > l2 {
			return getKth(arr2, start2, end2, arr1, start1, end1, k)
		}
		if l1 == 0 {
			return arr2[start2+k-1]
		}
		if k == 1 {
			return min(arr1[start1], arr2[start2])
		}
		i := start1 + min(l1, (k>>1)) - 1
		j := start2 + min(l2, (k>>1)) - 1

		if arr1[i] < arr2[j] {
			return getKth(arr1, i+1, end1, arr2, start2, end2, k-(i-start1+1))
		}
		return getKth(arr1, start1, end1, arr2, j+1, end2, k-(j-start2+1))

	}

	n, m := len(nums1), len(nums2)
	k1, k2 := (n+m+1)>>1, (n+m+2)>>1

	return float64((getKth(nums1, 0, n-1, nums2, 0, m-1, k1) + getKth(nums1, 0, n-1, nums2, 0, m-1, k2))) / 2

}

//leetcode submit region end(Prohibit modification and deletion)
