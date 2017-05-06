package main

import (
	"fmt"
	"github.com/liuzl/newsmth/downloader"
)

func main() {
	requestInfo := &downloader.RequestInfo{
		Url:      "http://m.newsmth.net",
		Method:   "GET",
		UseProxy: false,
		Platform: "mobile",
	}

	responseInfo := downloader.Download(requestInfo)
	fmt.Println(responseInfo)
	fmt.Println(responseInfo.Content)
	fmt.Println(responseInfo.Text)
}
