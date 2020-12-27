package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

func main() {
	var g errgroup.Group
	var urls = []string{
		"http://www.baidu.com/",
		"http://www.baidu.com/",
		"http://www.1234567.com/",//假的
	}
	for _, url := range urls {
		// Launch a goroutine to fetch the URL.
		url := url
		g.Go(func() error {
			// Fetch the URL.
			resp, err := http.Get(url)
			if err == nil { // 这里记得关掉
				resp.Body.Close()
			}
			return err
		})
	}
	// Wait for all HTTP fetches to complete.
	err := g.Wait();
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}
	fmt.Println("Successfully fetched all URLs.")
}