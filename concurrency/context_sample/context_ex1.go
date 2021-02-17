package context_sample

import (
	"context"
	"fmt"
	"time"
)

func A(ctx context.Context, name string) {
	go B(ctx, name) //A调用了B
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "A退出")
			return
		default:
			fmt.Println(name, "A do something")
			time.Sleep(2 * time.Second)
		}
	}
}

func B(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "B退出")
			return
		default:
			fmt.Println(name, "B do something")
			time.Sleep(2 * time.Second)
		}
	}
}

func Exit_test() {
	ctx, cancel := context.WithCancel(context.Background())

	go A(ctx, "【请求1】") //模拟client来了1个连接请求

	time.Sleep(3 * time.Second)
	fmt.Println("client断开连接，通知对应处理client请求的A,B退出")
	cancel() //假设满足某条件client断开了连接，那么就传播取消信号，ctx.Done()中得到取消信号

	time.Sleep(3 * time.Second)

}
