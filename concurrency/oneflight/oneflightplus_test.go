package oneflight

import (
	"runtime/debug"
	"sync/atomic"
	"testing"
	"time"
)

func TestPanicDo2(t *testing.T) {
	var g Group2
	fn := func() (interface{}, error) {
		panic("invalid memory address or nil pointer deference")
	}

	const n = 5
	waited := int32(n)
	panicCount := int32(0)
	done := make(chan struct{})
	t.Log("waited", waited)

	for i := 0; i < n; i++ {
		go func() {
			defer func() {
				if err := recover(); err != nil {
					// t.Logf("Got panic : %v\n%s", err, debug.Stack())
					atomic.AddInt32(&panicCount, 1)
				}
				if atomic.AddInt32(&waited, -1) == 0 {
					t.Logf("close channel : %s\n", debug.Stack())
					close(done)
				}
			}()
			g.Do("key", fn)
		}()
	}

	select {
	case <-done:
		if panicCount != n {
			t.Errorf("Expected %d panic, but got %d", n, panicCount)
		}
	case <-time.After(time.Second):
		t.Fatalf("Do hanges")
	}
}
