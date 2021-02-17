package wait_group_sample

import (
	"fmt"
	"sync"
)

func Wg_ex1() {
	var wg sync.WaitGroup

	wg.Add(2) //添加需要完成的工作量2

	go func() {
		wg.Done() //完成工作量1
		fmt.Println("goroutine 1 完成工作！")
	}()

	go func() {
		wg.Done() //完成工作量1
		fmt.Println("goroutine 2 完成工作！")
	}()

	wg.Wait() //等待工作量2均完成
	fmt.Println("所有的goroutine均已完成工作！")
}
