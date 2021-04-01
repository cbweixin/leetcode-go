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

type Engine struct{}

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
	log.Fatal(http.ListenAndServe(":9999", engine))
}
