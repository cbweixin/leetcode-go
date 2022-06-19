package main

import "./concurrency/ratelimiter"

const MAXREQS = 50

var sem = make(chan int, MAXREQS)

type Request3 struct {
	a, b   int
	replyc chan int
}

func process(r *Request3) {
	// do something
}

func handle(r *Request3) {
	sem <- 1 // doesn't matter what we put in it
	process(r)
	<-sem // one empty place in the buffer: the next request can start
}

func server4(service chan *Request3) {
	for {
		request := <-service
		go handle(request)
	}
}

func main() {

	// service := make(chan *Request3)
	// go server4(service)
	ratelimiter.Serve()

}
