package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"sync/atomic"
)

func main() {
	// 禁用GC，并保证在main函数执行结束前恢复GC。
	defer debug.SetGCPercent(debug.SetGCPercent(-1))
	var count int32
	newFunc := func() interface{} {
		return atomic.AddInt32(&count, 1)
	}

	pool := sync.Pool{New: newFunc}

	v1 := pool.Get()
	fmt.Printf("Value 1: %v\n", v1)

	pool.Put(10)
	pool.Put(11)
	pool.Put(12)

	v2 := pool.Get()
	fmt.Printf("value 2 : %v\n", v2)

	// trigger gc, but when I ran the code, it actually didn't trigger gc, so it is not guranteed
	debug.SetGCPercent(100)
	runtime.GC()
	v3 := pool.Get()
	fmt.Printf("value 3 : %v\n", v3)
	pool.New = nil
	v4 := pool.Get()
	fmt.Printf("value 4 : %v\n", v4)

}
