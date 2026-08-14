package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and a slice of URLs found on that page
	Fetch(url string) (body string, urls []string, err error)
}

// SafeCache holds visited URLs protected by a mutex
type SafeCache struct {
	mu      sync.Mutex
	visited map[string]bool
}

// Visited returns true if the url has already been crawled, or marks it as visited and returns false if it's new
func (c *SafeCache) Visited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.visited[url] {
		return true
	}
	c.visited[url] = true
	return false
}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth
func Crawl(url string, depth int, fetcher Fetcher) {
	cache := &SafeCache{
		visited: make(map[string]bool),
	}

	var crawler func(string, int)
	var wg sync.WaitGroup

	crawler = func(u string, d int) {
		defer wg.Done()

		if d <= 0 {
			return
		}

		// Check if we already crawled this URL
		if cache.Visited(u) {
			return
		}

		body, urls, err := fetcher.Fetch(u)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", u, body)

		// Spawn worker goroutines for each child URL
		for _, childURL := range urls {
			wg.Add(1)
			go crawler(childURL, d-1)
		}
	}

	// Start the initial fetch
	wg.Add(1)
	go crawler(url, depth)

	// Wait for all spawned goroutines to finish
	wg.Wait()
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
