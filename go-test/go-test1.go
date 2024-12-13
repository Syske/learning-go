package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, World!")
	data, err := ioutil.ReadFile("C:\\Users\\syske\\Downloads\\dev-tt_isv-api-669c46d8-fp79b_isv-api.log")
	if err != nil {
		fmt.Println("file reading error", err)
		return
	}
	fmt.Println("contents of file:", string(data))
}
