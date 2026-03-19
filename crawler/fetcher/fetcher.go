package fetcher

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func Fetch(targetUrl string, client *http.Client) ([]string, error) {

	baseUrl, err := url.Parse(targetUrl)
	if err != nil {
		return nil, fmt.Errorf("(ERR) >> unexpected base url: %s", err)
	}

	resp, err := http.Get(baseUrl.String())
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
						rawURL := attr.Val

						parsedURL, err := url.Parse(rawURL)
						if err != nil {
							continue
						}

						absoluteURL := baseUrl.ResolveReference(parsedURL).String()

						if parsedURL.Scheme != "" && parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
							continue
						}

						links = append(links, absoluteURL)
						break
					}
				}
			}
		}
	}

	return links, nil
}
