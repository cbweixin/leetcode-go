package main

import (
	"fmt"
	"math"
)

type Square2 struct {
	side float32
}

type Circle2 struct {
	radius float32
}

type Shaper2 interface {
	Area2() float32
}

func (sq *Square2) Area2() float32 {
	return sq.side * sq.side
}

func (ci *Circle2) Area2() float32 {
	return ci.radius * ci.radius * math.Pi
}

func main() {
	var areaIntf Shaper2
	sq1 := new(Square2)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	// 如果忽略 areaIntf.(*Square) 中的 * 号，会导致编译错误：impossible type assertion: Square does not implement Shaper
	// (Area method has pointer receiver)。
	if t, ok := areaIntf.(*Square2); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle2); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

	var areaIntf2 Shaper2
	circle1 := new(Circle2)
	circle1.radius = 5

	areaIntf2 = circle1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf2.(*Square2); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	if u, ok := areaIntf2.(*Circle2); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}
