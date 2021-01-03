package main

import (
	"fmt"
	"time"
)

type TwoInts struct {
	a int
	b int
}

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

type myTime struct {
	time.Time //anonymous field
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

type B2 struct {
	thing int
}

// 如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法。否则，就在普通的值类型上定义方法。
func (b *B2) change() {
	b.thing = 1
}

// 按值传递，实例的拷贝，接收者的数据并未改变
func (b B2) change2() {
	b.thing = 2
}

func (b B2) write() string { return fmt.Sprint(b) }

func main() {

	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())

	fmt.Println(IntVector{1, 2, 3}.Sum())

	m := myTime{time.Now()}
	// 调用匿名Time上的String方法
	fmt.Println("Full time now:", m.String())
	// 调用myTime.first3Chars
	fmt.Println("First 3 chars:", m.first3Chars())

	// 注意 Go 为我们做了探测工作，我们自己并没有指出是否在指针上调用方法，Go 替我们做了这些事情。b1 是值而 b2 是指针，方法都支持运行了。
	var b1 B2     // b1是值
	b2 := new(B2) // b2是指针
	b1.change()
	fmt.Println(b1.write())
	b1.change2()
	fmt.Println(b1.write())

	b2.change()
	fmt.Println(b2.write())
	b2.change2()
	fmt.Println(b2.write())

}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}
