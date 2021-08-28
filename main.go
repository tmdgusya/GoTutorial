package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed = errors.New("Request Failed")

func main() {
	urls := []string {
		"https://www.naver.com/",
	}

	for _, url := range urls {
		resp, _ := hitURL(url)
		fmt.Println("Response", resp)
	}
}

func hitURL(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400{
		return nil, errRequestFailed
	}
	return resp, nil
} 