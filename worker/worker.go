package worker

import (
	"log"
	"net/http"
	"webcrawler/crawler"
	"webcrawler/crawler/fetcher"
)

func Worker(jobs <-chan crawler.Task, result chan<- crawler.Result, client *http.Client) {

	for j := range jobs {

		links, err := fetcher.Fetch(j.Url, client)
		if err != nil {
			log.Printf("(ERR) >> failed to fetch [%s]: %s", j.Url, err)
		}

		res := crawler.Result{
			Url:   j.Url,
			Links: links,
			Depth: j.CurrDepth,
			Error: err,
		}

		result <- res
	}
}
