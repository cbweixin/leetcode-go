package loadgen

import (
	"./lib"
	"bytes"
	"errors"
	"fmt"
	"strings"
	"time"
)

type ParamSet struct {
	Caller     lib.Caller    // 调用器。
	TimeoutNS  time.Duration // 响应超时时间，单位：纳秒。
	LPS        uint32        // 每秒载荷量。
	DurationNS time.Duration // 负载持续时间，单位：纳秒。
	// why ResultCh is * pointer type? why not "chan lib.CallResult"?
	// https://stackoverflow.com/questions/28447297/how-to-check-for-an-empty-struct
	// 由于结构体类型的零值不是nil ，因此如果 这个通道的元素类型是lib.CallResult 的话，就会给后面对其中元素 值的零值判断带来小麻烦。
	// 我使用它的指针类型作为通道的元素类型，既 可以消除麻烦，也可以减少元素值复制带来的开销
	ResultCh chan *lib.CallResult // 调用结果通道。
}

// Check 会检查当前值的所有字段的有效性。
// 若存在无效字段则返回值非nil。
func (pset *ParamSet) Check() error {
	var errMsgs []string

	if pset.Caller == nil {
		errMsgs = append(errMsgs, "Invalid caller!")
	}
	if pset.LPS == 0 {
		errMsgs = append(errMsgs, "Invalid lps(load per second)!")
	}
	if pset.TimeoutNS == 0 {
		errMsgs = append(errMsgs, "Invalid timeoutNS!")
	}
	if pset.DurationNS == 0 {
		errMsgs = append(errMsgs, "Invalid durationNS!")
	}
	if pset.ResultCh == nil {
		errMsgs = append(errMsgs, "Invalid result channel!")
	}

	var buf bytes.Buffer
	buf.WriteString("Checking the parameters...")
	if errMsgs != nil {
		errMsg := strings.Join(errMsgs, " ")
		buf.WriteString(fmt.Sprintf("NOT passed! (%s)", errMsg))
		logger.Infoln(buf.String())
		return errors.New(errMsg)
	}
	buf.WriteString(
		fmt.Sprintf("Passed. (timeoutNS=%s, lps=%d,durationNS=%s)",
			pset.TimeoutNS, pset.LPS, pset.DurationNS))
	logger.Infoln(buf.String())
	return nil
}
