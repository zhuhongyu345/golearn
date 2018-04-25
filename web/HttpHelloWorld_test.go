package web

import (
	"net/http"
	"fmt"
	"strings"
	"testing"
	"log"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Form)
	fmt.Println(r.URL)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello World!")
}

func TestHW(t *testing.T) {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(":9090", nil)
	fmt.Println("start")
	if err != nil {
		log.Fatal("error:", err)
	}
}
