package lib

import "time"

type Caller interface {
	BuildReq() RawReq
	Call(req []byte, timeoutNs time.Duration) ([]byte, error)
	CheckResp(rawReq RawReq, rawResp RawResp) *CallResult
}
