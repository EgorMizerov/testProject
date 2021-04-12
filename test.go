package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	count := 10000
	t := time.Now()

	wg := sync.WaitGroup{}
	wg.Add(count)

	get := func() {
		defer wg.Done()
		http.Get("http://localhost:8010/json/hackers")
	}

	for i := 0; i < count; i++ {
		go get()
	}

	wg.Wait()
	fmt.Println(time.Since(t))
}
