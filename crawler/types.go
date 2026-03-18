package crawler

type Task struct {
	Url       string
	CurrDepth int
}

type Result struct {
	Url    string
	Links  []string
	Images []string
	Error  error
}
