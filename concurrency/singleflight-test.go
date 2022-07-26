package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/sync/singleflight"
)

var count int32

func getArticle(id int) (article string, err error) {
	atomic.AddInt32(&count, 1)
	time.Sleep(time.Duration(count) * time.Millisecond)

	return fmt.Sprintf("article: %d", id), nil
}

func singleflightGetArticle(sg *singleflight.Group, id int) (string, error) {
	v, err, _ := sg.Do(
		fmt.Sprintf("%d", id), func() (interface{}, error) {
			return getArticle(id)
		},
	)
	return v.(string), err
}

func main() {
	time.AfterFunc(
		2000*time.Millisecond, func() {
			atomic.AddInt32(&count, -count)
		},
	)

	fmt.Println("count", count)

	var (
		wg  sync.WaitGroup
		now = time.Now()
		n   = 1000
		sg  = &singleflight.Group{}
	)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			res, _ := singleflightGetArticle(sg, 1)
			if res != "article: 1" {
				panic("err")
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("send %d requests concurently, time used: %s\n", n, time.Since(now))

	now = time.Now()
	fmt.Println("count", count)

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func() {
			// res, _ := singleflightGetArticle(sg, 1)
			res, _ := getArticle(1)
			if res != "article: 1" {
				panic("err")
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Printf("send %d requests concurently, time used: %s\n", n, time.Since(now))
	fmt.Println("count", count)
}
