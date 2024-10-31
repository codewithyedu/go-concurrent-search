package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Simulated search replicas for web, image, and video searches
var (
	web1 = fakeSearch("web")
	web2 = fakeSearch("web")
	web3 = fakeSearch("web")

	image1 = fakeSearch("image")
	image2 = fakeSearch("image")
	image3 = fakeSearch("image")

	video1 = fakeSearch("video")
	video2 = fakeSearch("video")
	video3 = fakeSearch("video")
)

type Result string

type Search func(query string) Result

func main() {
	query := "golang"

	start := time.Now()
	results := webSearch(query)
	elapsed := time.Since(start)

	for _, result := range results {
		fmt.Println(result)
	}
	fmt.Println("execution time:", elapsed)
}

// fakeSearch simulates a search by delaying the result with random latency
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%v result for %v", kind, query))
	}
}

// webSearch initiates concurrent searches and returns the first result for each category
func webSearch(query string) (results []Result) {
	c := make(chan Result)
	timeout := time.After(80 * time.Millisecond)

	// Launch concurrent searches for each category (web, image, video)
	go func() { c <- getFirstResponse(query, web1, web2, web3) }()
	go func() { c <- getFirstResponse(query, image1, image2, image3) }()
	go func() { c <- getFirstResponse(query, video1, video2, video3) }()

	// Collect results or exit if a timeout occurs
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("--- timeout ---")
			return results
		}
	}

	return results
}

// getFirstResponse returns the first result from multiple replicas
func getFirstResponse(query string, replicas ...Search) Result {
	c := make(chan Result)

	// Start a search on each replica concurrently
	for i := range replicas {
		go func(i int) { c <- replicas[i](query) }(i)
	}

	return <-c
}
