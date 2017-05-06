package downloader

import (
	"testing"
)

func TestDownload(t *testing.T) {
	requestInfo := &RequestInfo{
		Url:      "http://m.newsmth.net",
		Method:   "GET",
		UseProxy: false,
		Platform: "mobile",
	}

	responseInfo := Download(requestInfo)
	if responseInfo.Error != nil {
		t.Error(responseInfo.Error)
	}
}
