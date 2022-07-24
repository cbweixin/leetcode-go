package main

import (
	"fmt"
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

}
