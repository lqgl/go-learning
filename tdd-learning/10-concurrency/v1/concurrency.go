package concurrency

type WebsiteChecker func(url string) bool

func CheckWebSites(wc WebsiteChecker, urls []string) map[string]bool {
	result := make(map[string]bool, 0)
	for _, url := range urls {
		result[url] = wc(url)
	}
	return result
}
