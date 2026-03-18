package config

import (
	"flag"
	"fmt"
	"net/url"
)

type Config struct {
	Url     string
	Depth   int
	Workers int
}

func ParseFlags() (*Config, error) {

	targetUrl := flag.String("url", "", "target url")
	depth := flag.Int("depth", 1, "maximum crawling depth")
	workers := flag.Int("workers", 5, "maximum count of workers")

	flag.Parse()

	if *targetUrl == "" {
		return nil, fmt.Errorf("(ERR) >> invlid input: ")
	}

	if _, err := url.ParseRequestURI(*targetUrl); err != nil {
		return nil, fmt.Errorf("(ERR) >> invlid input: %s", err)
	}

	config := Config{
		Url:     *targetUrl,
		Depth:   *depth,
		Workers: *workers}

	return &config, nil
}
