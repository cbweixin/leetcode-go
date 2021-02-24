package loadgen

import (
	helper "./testhelper"
	"testing"
)

var printDetail = false

func TestStart(t *testing.T) {
	server := helper.NewTCPServer()
	defer server.Close()

	serverAddr := "127.0.0.1:8080"
	t.Logf("Startup TCP server(%s)...\n", serverAddr)

}
