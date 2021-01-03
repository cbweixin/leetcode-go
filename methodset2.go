package main

import "fmt"

// 在 lst 上调用 CountInto 时会导致一个编译器错误，因为 CountInto 需要一个 Appender，而它的方法 Append 只定义在指针上。
// 在 lst 上调用 LongEnough 是可以的，因为 Len 定义在值上。
//
// 在 plst 上调用 CountInto 是可以的，因为 CountInto 需要一个 Appender，并且它的方法 Append 定义在指针上。 在 plst 上调用
// LongEnough 也是可以的，因为指针会被自动解引用。

// Go 语言规范定义了接口方法集的调用规则：
//
//类型 T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
//类型 *T 的可调用方法集包含接受者为 *T 的所有方法
//类型 *T 的可调用方法集不包含接受者为 T 的方法

type List2 []int

func (l List2) Len() int {
	return len(l)
}

func (l *List2) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func main() {
	// A bare value
	var lst List2
	// compiler error:
	// cannot use lst (type List) as type Appender in argument to CountInto:
	//       List does not implement Appender (Append method has pointer receiver)
	//CountInto(lst, 1, 10)
	if LongEnough(lst) { // VALID:Identical receiver type
		fmt.Printf("- lst is long enough\n")
	}

	// A pointer value
	plst := new(List2)
	CountInto(plst, 1, 10) //VALID:Identical receiver type
	if LongEnough(plst) {
		// VALID: a *List can be dereferenced for the receiver
		fmt.Printf("- plst is long enough\n")
	}

}
