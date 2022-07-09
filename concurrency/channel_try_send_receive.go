package main

import "fmt"

func main() {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: // go here if c is full
		}
	}

	tryReceieve := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-" // go here if c is empty
		}
	}

	trySend("Hello!") // succeed to send
	trySend("Hi!")    // succeed to send
	// Fail to send, but will not block.
	trySend("Bye!")
	// The following two lines will
	// both succeed to receive.
	fmt.Println(tryReceieve()) // Hello!
	fmt.Println(tryReceieve()) // Hi!
	// The following line fails to receive.
	fmt.Println(tryReceieve()) // -

	// below code has 50% possibility to panic, both of the two `case` operations are non-block
	c2 := make(chan struct{})
	close(c2)
	select {
	// send to closed, non-nil channel would cause panic
	case c2 <- struct{}{}:
		// but receive from closed, non-nil chan, you would get zero value, no panic
	case <-c2:
	}
}
