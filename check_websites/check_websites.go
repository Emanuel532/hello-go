package main

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, websiteUrls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range websiteUrls {
		go func() {
			//results[url] = wc(url)
			resultsChannel <- result{url, wc(url)}
		}()
	}
	for range websiteUrls {
		r := <-resultsChannel
		results[r.string] = r.bool
	}
	return results
}
