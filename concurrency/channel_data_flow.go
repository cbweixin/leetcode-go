package main

import (
	"encoding/binary"
	"math/rand"
	"sync"
)

func RandomGenerator() <-chan uint64 {
	c := make(chan uint64)
	go func() {
		rnds := make([]byte, 8)
		for {
			_, err := rand.Read(rnds)
			if err != nil {
				close(c)
				break
			}
			c <- binary.BigEndian.Uint64(rnds)
		}
	}()

	return c
}

// Data aggregation

// A data aggregation module worker aggregates several data streams of the same data type into one stream. Assume
// the data type is int64, then the following function will aggregate an arbitrary number of data streams into one.

func Aggregator(inputs ...<-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	for _, in := range inputs {
		go func(in <-chan uint64) {
			for {
				out <- <-in // <=> out <- (<-in)
			}
		}(in)
	}

	return out
}

// A better implementation should consider whether or not an input stream has been closed. (Also valid for the following
// other module worker implementations.)
func AggregatorImproved(inputs ...<-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Add(1)
		go func(in <-chan uint64) {
			defer wg.Done()
			// if in is closed, then the loop will ends eventually
			for x := range in {
				out <- x
			}

		}(in)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

// If the number of aggregated data streams is very small (two or three),
// we can use select block to aggregate ithese data streams.
func AggregatorWithTwoStream(inputs ...<-chan uint64) <-chan uint64 {
	out := make(chan uint64)
	go func() {
		inA, inB := inputs[0], inputs[1]
		for {
			select {
			case v := <-inA:
				out <- v
			case v := <-inB:
				out <- v
			}
		}
	}()

	return out
}

// Data division

// A data division module worker does the opposite of a data aggregation module worker. It is easy to implement a
// division worker, but in practice, division workers are not very useful and seldom used.
func Divisor(input <-chan uint64, outputs ...chan<- uint64) {
	for _, out := range outputs {
		go func(o chan<- uint64) {
			for {
				o <- <-input
			}
		}(out)
	}
}

// Data duplication/proliferation

// Data duplication (proliferation) can be viewed as special data decompositions. One piece of data will be duplicated
// and each of the duplicated data will be sent to different output data streams.
func Duplicator(in <-chan uint64) (<-chan uint64, <-chan uint64) {
	outA, outB := make(chan uint64), make(chan uint64)
	go func() {
		for x := range in {
			outA <- x
			outB <- x
		}
	}()

	return outA, outB
}
