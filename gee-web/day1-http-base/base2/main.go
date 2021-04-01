package main

import (
	"fmt"
	"log"
	"net/http"
)

// ❯ curl localhost:9999/
//url path = "/"
//❯ curl localhost:9999/hello
//header ["User-Agent"] = ["curl/7.64.1"]
//header ["Accept"] = ["*/*"]
//❯ curl localhost:9999/hellow
//404 NOT FOUND /hellow

type Engine struct {
	// 我们定义了一个空的结构体Engine，实现了方法ServeHTTP。这个方法有2个参数，第二个参数是 Request ，该对象包含了该HTTP请求的所有的信息，
	//比如请求地址、Header和Body等信息；第一个参数是 ResponseWriter ，利用 ResponseWriter 可以构造针对该请求的响应。
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "url path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "header [%q] = %q\n", k, v)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUND %s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	// 第二个参数的类型是什么呢？通过查看net/http的源码可以发现，Handler是一个接口，需要实现方法 ServeHTTP ，也就是说，只要传入任何实现了
	//ServerHTTP 接口的实例，所有的HTTP请求，就都交给了该实例处理了。
	log.Fatal(http.ListenAndServe(":9999", engine))
}
