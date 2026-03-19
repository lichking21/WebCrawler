package main

import (
	"net/http"
	"webcrawler/crawler"
	"webcrawler/worker"
)

type Pool struct {
	Client   http.Client
	Workers  int
	Jobs     int
	Seed     string
	MaxDepth int
}

func NewPool(client *http.Client, workersCount int, jobsCount int, seed string, maxDepth int) *Pool {
	return &Pool{
		Client:   *client,
		Workers:  workersCount,
		Jobs:     jobsCount,
		Seed:     seed,
		MaxDepth: maxDepth,
	}
}

func (p *Pool) StartPool() {

	cache := crawler.NewSafeSet()
	jobs := make(chan crawler.Task, p.Jobs)
	result := make(chan crawler.Result, p.Jobs)

	for w := 1; w <= p.Workers; w++ {
		go worker.Worker(jobs, result, &p.Client)
	}

	activeTask := 1
	cache.Add(p.Seed)

	jobs <- crawler.Task{Url: p.Seed, CurrDepth: 0}

	for res := range result {

		activeTask--

		for _, link := range res.Links {

			if res.Depth < p.MaxDepth && cache.Add(link) {

				newTask := crawler.Task{Url: link, CurrDepth: res.Depth + 1}
				activeTask++

				go func(t crawler.Task) {
					jobs <- t
				}(newTask)
			}
		}

		if activeTask == 0 {
			close(jobs)
			return
		}
	}
}
