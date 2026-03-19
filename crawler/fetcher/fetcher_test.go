package fetcher

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetch(t *testing.T) {
	client := http.Client{}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`
			<html>
				<body>
					<a href="http://link1.com">Link 1</a>
					<a href="/relative-link">Link 2</a>
				</body>
			</html>
		`))
	}))

	defer mockServer.Close()

	links, err := Fetch(mockServer.URL, &client)

	if err != nil {
		t.Errorf("(TEST) >> unexpected error: %s", err)
	}
	if len(links) != 2 {
		t.Errorf("(TEST) >> expected 2 links, recived: %s", err)
	}
	if links[0] != "http://link1.com" {
		t.Error("(TEST) >> failed to parse 1st link")
	}
	if links[1] != "/relative-link" {
		t.Error("(TEST) >> failed to parse 2nd link")
	}
}
