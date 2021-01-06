package main

import "fmt"

type Request struct {
	a, b   int
	replyc chan int // reply channel inside the Request
}

type binOp func(a, b int) int

func run(op binOp, req *Request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *Request) {
	for {
		req := <-service // requests arrive here
		// start goroutine for request:
		go run(op, req) // don't wait for op
	}
}

func server2(op binOp, service chan *Request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func startServer(op binOp) chan *Request {
	reqChan := make(chan *Request)
	go server(op, reqChan)
	return reqChan
}

// 在上一个版本中server在main函数返回后并没有完全关闭，而被强制结束了。为了改进这一点，我们可以提供一个退出通道给server
func startServer2(op binOp) (service chan *Request, quit chan bool) {
	service = make(chan *Request)
	quit = make(chan bool)
	go server2(op, service, quit)
	return service, quit
}

func main() {
	adder := startServer(func(a, b int) int { return a + b })
	const N = 100
	var reqs [N]Request
	for i := 0; i < N; i++ {
		req := &reqs[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder <- req
	}
	// checks:
	for i := N - 1; i >= 0; i-- { // doesn't matter what order
		if <-reqs[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}
	fmt.Println("done")

	adder2, quit := startServer2(func(a, b int) int { return a + b })
	var reqs2 [N]Request
	for i := 0; i < N; i++ {
		req := &reqs2[i]
		req.a = i
		req.b = i + N
		req.replyc = make(chan int)
		adder2 <- req
	}
	// checks:
	for i := N - 1; i >= 0; i-- { // doesn't matter what order
		if <-reqs2[i].replyc != N+2*i {
			fmt.Println("fail at", i)
		} else {
			fmt.Println("Request ", i, " is ok!")
		}
	}

	quit <- true
	fmt.Println("done")
}
