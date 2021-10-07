package main

import "sync"

type URLChecker func(string) bool

// avoid concurrent map writes by using sync.Map
func CheckURL(checker URLChecker, urls []string) (m sync.Map) {
	// run each url checker concurrently
	wg := &sync.WaitGroup{}
	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			ok := checker(url)
			m.Store(url, ok)
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
