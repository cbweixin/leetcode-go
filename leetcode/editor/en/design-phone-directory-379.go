package main

type PhoneDirectory1 struct {
	Availables   *Node
	NoAvailables *Node
}

type Node struct {
	Number int
	Next   *Node
}

func Constructor1(maxNumbers int) PhoneDirectory1 {
	root := &Node{Number: 0}
	node := root
	for i := 1; i < maxNumbers; i++ {
		node.Next = &Node{Number: i}
		node = node.Next
	}
	return PhoneDirectory1{Availables: root}
}

func (this *PhoneDirectory1) Get1() int {
	if this.Availables == nil {
		return -1
	}
	node := this.Availables
	this.Availables = this.Availables.Next
	node.Next = this.NoAvailables
	this.NoAvailables = node
	return this.NoAvailables.Number
}

func (this *PhoneDirectory1) Check1(number int) bool {
	node := this.Availables
	for node != nil {
		if node.Number == number {
			return true
		}
		node = node.Next
	}
	return false
}

func (this *PhoneDirectory1) Release1(number int) {
	var prev *Node
	node := this.NoAvailables
	for node != nil {
		if node.Number == number {
			if prev == nil {
				this.NoAvailables = node.Next
			} else {
				prev.Next = node.Next
			}
			node.Next = this.Availables
			this.Availables = node
			return
		}
		prev = node
		node = node.Next
	}
}

type PhoneDirectory struct {
	dir  []bool
	free []int
}

/** Initialize your data structure here
  @param maxNumbers - The maximum numbers that can be stored in the phone directory. */
func Constructor(maxNumbers int) PhoneDirectory {
	free := make([]int, maxNumbers)
	for i := range free {
		free[i] = len(free) - 1 - i
	}
	return PhoneDirectory{make([]bool, maxNumbers), free}
}

/** Provide a number which is not assigned to anyone.
  @return - Return an available number. Return -1 if none is available. */
func (this *PhoneDirectory) Get() int {
	if len(this.free) == 0 {
		return -1
	}
	n := this.free[len(this.free)-1]
	this.free = this.free[:len(this.free)-1]
	this.dir[n] = true
	return n
}

/** Check if a number is available or not. */
func (this *PhoneDirectory) Check(number int) bool {
	return !this.dir[number]
}

/** Recycle or release a number. */
func (this *PhoneDirectory) Release(number int) {
	if this.dir[number] {
		this.dir[number] = false
		this.free = append(this.free, number)
	}
}

func main() {

}
