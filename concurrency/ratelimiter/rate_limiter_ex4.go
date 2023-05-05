package ratelimiter

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"
)

// https://www.bbsmax.com/A/gGdXeaAvz4/
// 看了两篇关于golang中限流器的帖子：
//
//    Gin 开发实践：如何实现限流中间件
//    常用限流策略——漏桶与令牌桶介绍
//
//我照着用，居然没效果……
//
//时间有限没有深究。这实在是一个很简单的功能，我的需求是每分钟限制大约xxx次请求，平均每秒限制到XXX除60次左右的请求也就够了。

// Limiter 限流器对象
type Limiter struct {
	value int64
	max   int64
	ts    int64
}

// NewLimiter 产生一个限流器
func NewLimiter(cnt int64) *Limiter {
	return &Limiter{
		value: 0,
		max:   cnt,
		ts:    time.Now().Unix(),
	}
}

// Ok 是否可以通过
func (l *Limiter) Ok() bool {
	// Unix() return seconds elapsed
	ts := time.Now().Unix()
	tsOld := atomic.LoadInt64(&l.ts)
	if ts != tsOld {
		atomic.StoreInt64(&l.ts, ts)
		atomic.StoreInt64(&l.value, 1)
		return true
	}
	return atomic.AddInt64(&(l.value), 1) < l.max
}

// SetMax 设置最大限制
func (l *Limiter) SetMax(m int64) {
	l.max = m
}

// MaxAllowed 限流器
func MaxAllowed(limitValue int64) func(c *gin.Context) {
	limiter := NewLimiter(limitValue)
	log.Println("limiter.SetMax:", limitValue)
	// 返回限流逻辑
	return func(c *gin.Context) {
		if !limiter.Ok() {
			c.AbortWithStatus(http.StatusServiceUnavailable) //超过每秒200，就返回503错误码
			return
		}
		c.Next()
	}
}
func main() {
	router := gin.New()
	router.Use(MaxAllowed(200)) //限制每秒最多允许200个请求
}
