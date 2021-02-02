package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {

	for i := 1; i < 1000; i++ {
		go func() {
			time.Sleep(1e2)
			resp, err := http.PostForm("http://localhost:8088/test2",
				url.Values{"in": {"barabcaadfadafd"}})

			if nil != err {
				fmt.Println("errorination happened getting the response", err)
				return
			}

			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)

			if nil != err {
				fmt.Println("errorination happened reading the body", err)
				return
			}

			fmt.Println(string(body[:]))

		}()
	}
	time.Sleep(10e9)
}
