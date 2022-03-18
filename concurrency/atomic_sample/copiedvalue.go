package atomic_sample

import (
	"fmt"
	"sync/atomic"
)

// https://www.bilibili.com/video/BV1NW41137jt?t=1.4
// https://www.bilibili.com/video/BV1n4411b75w?t=1470.3

func StoreValue() {
	var countVal atomic.Value
	countVal.Store([]int{1, 3, 5, 7})
	anotherStore(countVal)
	fmt.Printf("The count value: %+v \n", countVal.Load())
}

func anotherStore(countVal atomic.Value) {
	countVal.Store([]int{2, 4, 6, 8})
}
