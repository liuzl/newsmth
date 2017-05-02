package main

import (
	"fmt"
	"github.com/golang/glog"
	"github.com/mozillazg/request"
	"net"
	"net/http"
	"time"
)

func GetProxy() string {
	url := "http://127.0.0.1:8888/get"
	c := new(http.Client)
	req := request.NewRequest(c)
	req.Client.Timeout = time.Duration(3 * time.Second)
	for i := 0; i < 3; i++ {
		time.Sleep(1 * time.Second)
		resp, err := req.Get(url)
		if err == nil {
			body, readErr := resp.Text()
			resp.Body.Close()
			if readErr == nil {
				return fmt.Sprintf("http://%s", body)
			}
		} else {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				glog.Info("Get proxy timeout")
				continue
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(GetProxy())
}
