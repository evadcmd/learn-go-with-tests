package main

type URLChecker func(string) bool

func CheckURL(checker URLChecker, urls []string) map[string]bool {
	m := make(map[string]bool)
	for _, url := range urls {
		m[url] = checker(url)
	}
	return m
}

func dummyChecker(string) bool {
	return true
}

func main() {
	CheckURL(dummyChecker, []string{"a", "b", "c"})
}
