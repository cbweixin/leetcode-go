package main

import "fmt"

type Shaper3 interface {
	Area() float32
}

type TopologicalGenus interface {
	Rank() int
}

type Square3 struct {
	side float32
}

func (sq *Square3) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square3) Rank() int {
	return 1
}

type Rectangle2 struct {
	length, width float32
}

func (r Rectangle2) Area() float32 {
	return r.length * r.width
}

func (r Rectangle2) Rank() int {
	return 2
}

func main() {
	r := Rectangle2{5, 3} // Area() of Rectangle2 needs a value
	q := &Square3{5}      // Area() of Square3 needs a pointer
	shapes := []Shaper3{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
	topgen := []TopologicalGenus{r, q}
	fmt.Println("Looping through topgen for rank ...")
	for n, _ := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}

}
