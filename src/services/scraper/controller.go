package scraper

type Controller interface {
	FetchFromReddit(url string) []string
}
