package downloader

import (
	"http"
	"time"
)

func Download(requestInfo *RequestInfo) *ResponseInfo {
	var timeout time.Duration
	if requestInfo.Timeout > 0 {
		timeout = time.Duration(requestInfo.Timeout) * time.Second
	} else {
		timeout = 30 * time.Second
	}
	client := &http.Client{
		Timeout: timeout,
	}
	responseInfo := &ResponseInfo{
		Url: requestInfo.Url,
	}
	transport := http.Transport{
		DisableKeepAlives: true,
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
	}

	//proxy
	if requestInfo.UseProxy {
		var proxy string
		var err error

	}
	return nil
}
