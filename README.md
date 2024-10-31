# Go Concurrent Search

A simple Go program that demonstrates concurrent searching with timeouts. This code simulates a search engine querying multiple sources (web, image, and video) and returns the first response received from each category or exits if a timeout occurs.

## Features
- **Concurrent Search**: Launches multiple search queries in parallel across three categories: web, image, and video.
- **First Response Handling**: Returns the first response received for each category, allowing faster retrieval.
- **Timeouts**: If a search category does not return a result within the specified time, the program moves on without blocking.

## How It Works
The program:
1. Creates "fake" search functions (`fakeSearch`) with random response times to simulate real search responses.
2. Uses goroutines to execute searches for each category concurrently.
3. Captures the first available result per category or times out after a set duration (80ms).
4. Returns the results for each category or exits if the timeout is reached.

## Usage

To run this program, you'll need [Go installed](https://golang.org/doc/install).

Clone the repository and run the following commands:

```bash
go run main.go
```

You should see output resembling:

```bash
web result for golang
image result for golang
video result for golang
execution time: 83.28327ms
```

If any category times out, you will see a `--- timeout ---` message.

## Code Highlights
- **Concurrency**: Uses goroutines and channels to run and manage search queries concurrently.
- **Timeouts**: Demonstrates a practical example of Go’s `time.After` to set deadlines on concurrent tasks.

## Potential Enhancements
- **Adjustable Timeout**: Allow customization of timeout durations for each search category.
- **Additional Categories**: Add more simulated categories or types of searches to enhance functionality.
- **Real API Integration**: Replace `fakeSearch` with actual search APIs to retrieve live results.

## License
This program is licensed under the MIT License.
