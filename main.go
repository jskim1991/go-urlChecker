package main

import (
	"errors"
	"fmt"
	"net/http"
)

type httpGetResponse struct {
	url    string
	status string
}

var (
	errorRequestFailed = errors.New("Request Failed")
)

func main() {
	c := make(chan httpGetResponse)
	results := make(map[string]string) // same as map[string]string{}
	urls := []string{
		"https://naver.com",
		"https://jskim1991.medium.com",
		"https://google.com",
		"https://stackoverflow.com",
		"https://github.com",
	}

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < len(urls); i++ {
		httpGetResponse := <-c
		results[httpGetResponse.url] = httpGetResponse.status
	}

	for k, v := range results {
		fmt.Println(k, v)
	}
}

func hitURL(url string, c chan<- httpGetResponse) {
	response, error := http.Get(url)
	status := "OK"
	if error != nil || response.StatusCode >= 400 {
		status = "FAILED"
	}
	c <- httpGetResponse{url: url, status: status}
}
