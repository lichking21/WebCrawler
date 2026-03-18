package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webcrawler/crawler"
	"webcrawler/crawler/fetcher"
)

func main() {

	client := http.Client{Timeout: 10 * time.Second}
	cache := crawler.NewSafeSet()

	url := "https://go.dev"

	cache.Add(url)

	fmt.Printf("Fetching [%s]...\n", url)

	links, err := fetcher.Fetch(url, &client)
	if err != nil {
		log.Fatalf("(ERR) >> failed to fetch [%s]: %s", url, err)
	}

	for _, link := range links {
		if cache.Add(link) {
			fmt.Printf("[+] New url: %s\n", link)
		} else {
			fmt.Printf("[-] Duplikate: %s\n", link)
		}
	}
}
