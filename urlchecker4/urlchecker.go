package main

type URLChecker func(string) bool

type entry struct {
	url string
	ok  bool
}

func CheckURL(checker URLChecker, urls []string) map[string]bool {
	m := make(map[string]bool)
	ch := make(chan entry, 10)
	defer close(ch)
	for _, url := range urls {
		go func(url string) {
			ch <- entry{url, checker(url)}
		}(url)
	}
	for range urls {
		e := <-ch
		m[e.url] = e.ok
	}
	return m
}

func dummyChecker(string) bool {
	return true
}

func main() {
	CheckURL(dummyChecker, []string{"a", "b", "c"})
}
