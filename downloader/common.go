package downloader

import "fmt"

type RequestInfo struct {
	Url      string
	Method   string
	PostData string
	UseProxy bool
	Proxy    string
	Timeout  int
	MaxLen   int64
	Platform string
}

func (req *RequestInfo) String() string {
	return fmt.Sprintf("url: %s, post data: %s, proxy: %s", req.Url, req.PostData, req.Proxy)
}

type ResponseInfo struct {
	Url        string
	Text       string
	Content    []byte
	Encoding   string
	StatusCode int
	Proxy      string
	Cookies    map[string]string
	Error      *DownloaderError
}

func (resp *ResponseInfo) String() string {
	return fmt.Sprintf("url: %s, status: %d, proxy: %s, error: %s", resp.Url, resp.StatusCode, resp.Proxy, resp.Error)
}

type DownloaderError struct {
	Code    int
	Message string
}

func (e DownloaderError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
