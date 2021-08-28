package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	
	go goCount("roach")
	goCount("bbbb")	
	
}

func goCount(person string) {
	for i:= 0; i < 10; i++ {
		fmt.Println(person, "is Person", i)
		time.Sleep(time.Second)
	}
}

func makeURLArray() []string {
	return []string {
		"https://www.naver.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://www.daum.net/",
		"http://www.roach-devlog.com/",
	}
}

func printURLResp(urls []string) {
	results := make(map[string]string)

	for _, url := range urls {
		var result string
		resp, err := hitURL(url)
		if err != nil {
			result = "FAIL"
		}
		if resp != nil {
			result = "SUCCESS"
		}
		results[url] = result
	}
	for url, result := range results {
		fmt.Println(url, result)
	}
}

func hitURL(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400{
		return nil, errRequestFailed
	}
	return resp, nil
} 