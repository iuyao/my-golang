package main

import (
	"fmt"
	"log"
	"my-golang/go-web/web-base/base3/gee"
	"net/http"
)

func main() {
	engine := gee.NewEngine()
	engine.Get("/", handleIndex)
	engine.Get("/hello", handleHello)
	log.Fatal(engine.Run(":9999"))

}

func handleIndex(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
