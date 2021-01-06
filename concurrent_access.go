package main

import (
	"fmt"
	"strconv"
)

type Person3 struct {
	Name   string
	salary float64
	chF    chan func()
}

func NewPerson(name string, salary float64) *Person3 {
	p := &Person3{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person3) backend() {
	for f := range p.chF {
		f()
	}
}

// Set salary.
func (p *Person3) SetSalary(sal float64) {
	p.chF <- func() { p.salary = sal }
}

// Retrieve salary.
func (p *Person3) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
	return <-fChan
}

func (p *Person3) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func main() {
	// 为了保护对象被并发访问修改，我们可以使用协程在后台顺序执行匿名函数来替代使用同步互斥锁。在下面的程序中我们有一个类型Person其中包含
	// 一个字段chF，这是一个用于存放匿名函数的通道。
	//
	//这个结构在构造函数NewPerson()中初始化的同时会启动一个后台协程backend()。backend()方法会在一个无限循环中执行chF中放置的所有函数，
	// 有效的将它们序列化从而提供了安全的并发访问。更改和读取salary的方法会通过将一个匿名函数写入chF通道中，然后让backend()按顺序执行以
	// 达到其目的。需注意的是Salary方法创建的闭包函数是如何将fChan通道包含在其中的。

	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed:")
	fmt.Println(bs)

}
