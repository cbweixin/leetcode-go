package main

import (
	"fmt"
	"sync"
)

func main() {
	/**
	The lock is locked.
	Unlock the lock.
	The lock is unlocked
	Unlock the lock again
	fatal error: sync: unlock of unlocked mutex

	goroutine 1 [running]:
	runtime.throw({0x10a7e48, 0x1})
	        /Users/xin.wei/.asdf/installs/golang/1.17.8/go/src/runtime/panic.go:1198 +0x71 fp=0xc000106ea8 sp=0xc000106e78 pc=0x102f691
	sync.throw({0x10a7e48, 0x0})
	        /Users/xin.wei/.asdf/installs/golang/1.17.8/go/src/runtime/panic.go:1184 +0x1e fp=0xc000106ec8 sp=0xc000106ea8 pc=0x105737e
	sync.(*Mutex).unlockSlow(0xc000128008, 0xffffffff)
	        /Users/xin.wei/.asdf/installs/golang/1.17.8/go/src/sync/mutex.go:196 +0x3c fp=0xc000106ef0 sp=0xc000106ec8 pc=0x1065d1c
	sync.(*Mutex).Unlock(...)
	        /Users/xin.wei/.asdf/installs/golang/1.17.8/go/src/sync/mutex.go:190
	main.main()
	        /Users/xin.wei/github-personal/leetcode-go/concurrency/repeatly_unlock.go:22 +0x1a9 fp=0xc000106f80 sp=0xc000106ef0 pc=0x108baa9
	runtime.main()
	        /Users/xin.wei/.asdf/installs/golang/1.17.8/go/src/runtime/proc.go:255 +0x227 fp=0xc000106fe0 sp=0xc000106f80 pc=0x1031d27
	runtime.goexit()
	        /Users/xin.wei/.asdf/installs/golang/1.17.8/go/src/runtime/asm_amd64.s:1581 +0x1 fp=0xc000106fe8 sp=0xc000106fe0 pc=0x105b1c1


	*/
	defer func() {
		fmt.Println("Try to recover the panic.")
		if p := recover(); p != nil {
			fmt.Printf("Recovered panic (%#v).\n", p)
		}
	}()
	var mutex sync.Mutex
	mutex.Lock()
	fmt.Println("The lock is locked.")
	fmt.Println("Unlock the lock.")
	mutex.Unlock()
	fmt.Println("The lock is unlocked")
	fmt.Println("Unlock the lock again")
	mutex.Unlock()
}
