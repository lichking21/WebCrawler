package main

import (
	"log"
	"net/http"
	"time"
)

func main() {

	client := http.Client{Timeout: 10 * time.Second}
	//seed := "https://go.dev"
	seed := "http://stress-test.com"
	workersCount := 10
	jobsCount := 100
	maxDepth := 2

	pool := NewPool(&client, workersCount, jobsCount, seed, maxDepth)

	log.Printf("Starting Web Crawler Seed:[%s] | Workers:[%d] | MaxDepth:[%d]", seed, workersCount, maxDepth)

	pool.StartPool()

	log.Printf("All tasks done")
}
