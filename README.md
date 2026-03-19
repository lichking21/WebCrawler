# Go Concurrent Web Crawler

A high-performance, recursive web crawler built from scratch to demonstrate advanced concurrency patterns and state management in Golang. 

## 🚀 Core Architecture

* **Dynamic Worker Pool:** Implements a Fan-Out/Fan-In pattern using buffered channels for efficient HTTP fetching and task distribution without goroutine leaks.
* **Thread-Safe State:** Uses `sync.RWMutex` over a custom Set data structure to achieve `O(1)` URL deduplication and prevent infinite recursive loops.
* **Asynchronous Orchestrator:** A non-blocking event loop that tracks in-flight tasks and handles recursive job dispatching, completely eliminating circular deadlocks.
* **Stream Parsing:** Utilizes `golang.org/x/net/html` tokenizer for on-the-fly, low memory-footprint DOM parsing.

## 🛠 Tech Stack & Concepts
* **Language:** Go (Golang)
* **Concurrency:** Goroutines, Buffered Channels, Worker Pool
* **Synchronization:** `sync.RWMutex`
* **Networking:** `net/http` (Custom timeouts and header spoofing)

## ⚡ Quick Start
```bash
git clone https://github.com/lichking21/WebCrawler.git
cd go-crawler
go run .
