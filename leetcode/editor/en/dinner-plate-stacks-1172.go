package main

import "sort"

//You have an infinite number of stacks arranged in a row and numbered (left to
//right) from 0, each of the stacks has the same maximum capacity.
//
// Implement the DinnerPlates class:
//
//
// DinnerPlates(int capacity) Initializes the object with the maximum capacity o
//f the stacks.
// void push(int val) Pushes the given positive integer val into the leftmost st
//ack with size less than capacity.
// int pop() Returns the value at the top of the rightmost non-empty stack and r
//emoves it from that stack, and returns -1 if all stacks are empty.
// int popAtStack(int index) Returns the value at the top of the stack with the
//given index and removes it from that stack, and returns -1 if the stack with tha
//t given index is empty.
//
//
// Example:
//
//
//Input:
//["DinnerPlates","push","push","push","push","push","popAtStack","push","push",
//"popAtStack","popAtStack","pop","pop","pop","pop","pop"]
//[[2],[1],[2],[3],[4],[5],[0],[20],[21],[0],[2],[],[],[],[],[]]
//Output:
//[null,null,null,null,null,null,2,null,null,20,21,5,4,3,1,-1]
//
//Explanation:
//DinnerPlates D = DinnerPlates(2);  // Initialize with capacity = 2
//D.push(1);
//D.push(2);
//D.push(3);
//D.push(4);
//D.push(5);         // The stacks are now:  2  4
//                                           1  3  5
//                                           ﹈ ﹈ ﹈
//D.popAtStack(0);   // Returns 2.  The stacks are now:     4
//                                                       1  3  5
//                                                       ﹈ ﹈ ﹈
//D.push(20);        // The stacks are now: 20  4
//                                           1  3  5
//                                           ﹈ ﹈ ﹈
//D.push(21);        // The stacks are now: 20  4 21
//                                           1  3  5
//                                           ﹈ ﹈ ﹈
//D.popAtStack(0);   // Returns 20.  The stacks are now:     4 21
//                                                        1  3  5
//                                                        ﹈ ﹈ ﹈
//D.popAtStack(2);   // Returns 21.  The stacks are now:     4
//                                                        1  3  5
//                                                        ﹈ ﹈ ﹈
//D.pop()            // Returns 5.  The stacks are now:      4
//                                                        1  3
//                                                        ﹈ ﹈
//D.pop()            // Returns 4.  The stacks are now:   1  3
//                                                        ﹈ ﹈
//D.pop()            // Returns 3.  The stacks are now:   1
//                                                        ﹈
//D.pop()            // Returns 1.  There are no stacks.
//D.pop()            // Returns -1.  There are still no stacks.
//
//
//
// Constraints:
//
//
// 1 <= capacity <= 20000
// 1 <= val <= 20000
// 0 <= index <= 100000
// At most 200000 calls will be made to push, pop, and popAtStack.
//
// Related Topics Design
// 👍 217 👎 27

// 2021-01-20 12:24:56

//leetcode submit region begin(Prohibit modification and deletion)
type DinnerPlates struct {
	cap     int
	stacks  [][]int
	notFull []int
}

func Constructor(capacity int) DinnerPlates {
	return DinnerPlates{capacity, nil, nil}

}

func (this *DinnerPlates) Push(val int) {
	if len(this.notFull) > 0 {
		idx := this.notFull[0]
		this.notFull = this.notFull[1:]
		if idx < len(this.stacks) {
			this.stacks[idx] = append(this.stacks[idx], val)
			return
		}
		this.notFull = []int{}
	}
	if len(this.stacks) == 0 || len(this.stacks[len(this.stacks)-1]) == this.cap {
		this.stacks = append(this.stacks, []int{})
	}

	this.stacks[len(this.stacks)-1] = append(this.stacks[len(this.stacks)-1], val)

}

func (this *DinnerPlates) Pop() int {
	for len(this.stacks) > 0 {
		if len(this.stacks[len(this.stacks)-1]) == 0 {
			this.stacks = this.stacks[:len(this.stacks)-1]
		}
	}
	if len(this.stacks) == 0 {
		return -1
	}
	a := this.stacks[len(this.stacks)-1]
	m := a[len(a)-1]
	a = a[:len(a)-1]
	this.stacks[len(this.stacks)-1] = a
	return m

}

func (this *DinnerPlates) PopAtStack(index int) int {
	if index > len(this.stacks) || len(this.stacks) == 0 {
		return -1
	}
	l := len(this.notFull)
	if l == 0 {
		this.notFull = append(this.notFull, index)
	}
	// review, how to do bisect.insort(arr, ele) in go
	i := sort.Search(l, func(k int) bool { return this.notFull[k] > index })
	copy(this.notFull[i+1:], this.notFull[i:])
	this.notFull[i] = index

	m := this.stacks[len(this.stacks)-1]
	a := m[len(m)-1]
	m = m[:len(m)-1]

	this.stacks[len(this.stacks)] = m
	return a

}

/**
 * Your DinnerPlates object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * obj.Push(val);
 * param_2 := obj.Pop();
 * param_3 := obj.PopAtStack(index);
 */
//leetcode submit region end(Prohibit modification and deletion)
