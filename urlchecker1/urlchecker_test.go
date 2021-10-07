package main

import (
	"testing"
	"time"
)

const timeout = 3 * time.Second

// valid urls including its delay time
var validURLs = map[string]time.Duration{
	"http://tata": 250 * time.Nanosecond,
	"http://ruru": 520 * time.Nanosecond,
}

// if url in valid urls list then wait for its delay and return true,
// else wait for timeout and return false
func mockURLChecker(url string) bool {
	delay, ok := validURLs[url]
	if ok {
		time.Sleep(delay)
	} else {
		time.Sleep(timeout)
	}
	return ok
}

func TestURLChecker(t *testing.T) {
	urls := []string{
		"http://tata",
		"http://ruru",
		"http://unknown",
	}

	for url, got := range CheckURL(mockURLChecker, urls) {
		_, want := validURLs[url]
		if got != want {
			t.Errorf("url %s got %t want %t", url, got, want)
		}
	}

}
