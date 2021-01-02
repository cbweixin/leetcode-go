package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {

	// OK
	// 试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，但是 new() 一个 map 并试图向其填充数据，将会引发运行时错误！
	// 因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// NOT OK
	//z := make(Bar) // 编译错误：cannot make type Bar
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// NOT OK
	u := new(Foo)
	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	(*u)["y"] = "world"

}
