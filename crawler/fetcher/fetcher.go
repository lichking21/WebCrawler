package fetcher

import (
	"fmt"
	"log"
	"mime"
	"net/http"

	"golang.org/x/net/html"
)

func Fetch(targetUrl string, client *http.Client) ([]string, error) {

	resp, err := http.Get(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("(ERR) >> failed to get response: %s", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("(ERR) >> bad status code: %d", resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")

	// .ParseMediaType() is using to clear the line from charsets
	// Before parsing:"text/html; charset=utf-8", After parsing:"text/html"
	mediaType, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		return nil, fmt.Errorf("(ERR) >> failed to parse content type: %s", targetUrl)
	}

	if mediaType != "text/html" {
		log.Printf("(LOG) >> skipping %s: invalid type [%s]\n", targetUrl, mediaType)
		return nil, nil
	}

	var links []string
	z := html.NewTokenizer(resp.Body)

	for {

		tt := z.Next()

		if tt == html.ErrorToken {
			break
		}

		if tt == html.StartTagToken {

			t := z.Token()
			if t.Data == "a" {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
						break
					}
				}
			}
		}
	}

	return links, nil
}
