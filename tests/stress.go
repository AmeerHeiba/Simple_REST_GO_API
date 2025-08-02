package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func stressTest(name, url string, requests int) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	successes := 0
	start := time.Now()

	for i := 0; i < requests; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp, err := http.Get(url)
			if err == nil && resp.StatusCode == http.StatusOK {
				mu.Lock()
				successes++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("%s: %d/%d successful responses in %s\n", name, successes, requests, elapsed)
}

func main() {
	requests := 1000
	fmt.Println("Starting stress tests with", requests, "requests each...")
	stressTest("Standard Server (net/http)", "http://localhost:8080", requests)
	stressTest("Fiber Server", "http://localhost:8081", requests)
}
