package main

import "fmt"

type Request2 struct {
	a, b   int
	replyc chan int // reply channel inside the Request
}

type binOp2 func(a, b int) int

func run2(op binOp2, req *Request2) {
	req.replyc <- op(req.a, req.b)
}

func (r *Request2) String() string {
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.replyc)
}

func server3(op binOp2, service chan *Request2, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run2(op, req)
		case <-quit:
			return // stop infinite loop
		}
	}
}

func startServer3(op binOp2) (service chan *Request2, quit chan bool) {
	service = make(chan *Request2)
	quit = make(chan bool)
	go server3(op, service, quit)
	return service, quit
}

func main() {
	adder, quit := startServer3(func(a, b int) int { return a + b })
	// make requests:
	req1 := &Request2{3, 4, make(chan int)}
	req2 := &Request2{150, 250, make(chan int)}
	// send requests on the service channel
	adder <- req1
	adder <- req2
	// ask for the results: ( method String() is called )
	fmt.Println(req1, req2)
	// shutdown server:
	quit <- true
	fmt.Print("done")

}
