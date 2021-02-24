package loadgen

import (
	"./lib"
	helper "./testhelper"
	"testing"
	"time"
)

var printDetail = false

func TestStart(t *testing.T) {
	server := helper.NewTCPServer()
	defer server.Close()

	serverAddr := "127.0.0.1:8080"
	t.Logf("Startup TCP server(%s)...\n", serverAddr)
	err := server.Listen(serverAddr)
	if err != nil {
		t.Fatalf("TCP server startup failing! (addr=%s)", serverAddr)
		t.FailNow()
	}

	pset := ParamSet{
		Caller:     helper.NewTCPComm(serverAddr),
		TimeoutNS:  50 * time.Millisecond,
		LPS:        uint32(1000),
		DurationNS: 10 * time.Second,
		ResultCh:   make(chan *lib.CallResult, 50),
	}
	t.Logf("Initialize load generator (timeoutNS=%v, lps=%d, durationNS=%v)...",
		pset.TimeoutNS, pset.LPS, pset.DurationNS)
	gen, err := NewGenerator(pset)
	if err != nil {
		t.Fatalf("Load generator initialization failing: %s\n",
			err)
		t.FailNow()
	}

	t.Log("start load generator...")
	gen.Start()
}
