package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func Racer(ctx context.Context, urls ...string) (string, error) {
	ch := make(chan string)
	defer close(ch)
	for _, url := range urls {
		go func(url string) {
			http.Get(url)
			ch <- url
		}(url)
	}
	select {
	case first_url := <-ch:
		return first_url, nil
	case <-ctx.Done():
		return "", fmt.Errorf("timeout")
	}
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)
	fast_url, err := Racer(
		ctx,
		"https://cloud.google.com/bigquery/docs/quickstarts/quickstart-client-libraries#client-libraries-install-csharp",
		"https://www.envoyproxy.io/docs/envoy/v1.19.1/api-v3/config/listener/v3/listener.proto#envoy-v3-api-field-config-listener-v3-listener-connection-balance-config",
		"https://www.google.com/",
	)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(fast_url)
	}
}
