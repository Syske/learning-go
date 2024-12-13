package net_util

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func GetRequest(url string) string {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		result, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		return string(result)
	}
	return ""
}

func PostRequest(url string, contentType string, body string) string {
	fmt.Println(url)

	resp, err := http.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		result, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		return string(result)
	}
	return ""
}
