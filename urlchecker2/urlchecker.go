package main

import "sync"

type URLChecker func(string) bool

func CheckURL(checker URLChecker, urls []string) map[string]bool {
	m := make(map[string]bool)
	// run each url checker concurrently
	wg := &sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			m[url] = checker(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
	return m
}

func dummyChecker(string) bool {
	return true
}

func main() {
	CheckURL(dummyChecker, []string{"a", "b", "c"})
}
