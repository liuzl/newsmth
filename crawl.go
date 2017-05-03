package main

import (
	"errors"
	"fmt"
	"github.com/golang/glog"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/xpath"
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

func ParsePage(page []byte) (map[string]interface{}, error) {
	ret := make(map[string]interface{})
	doc, err := gokogiri.ParseHtml(page)
	if err != nil {
		fmt.Println(err)
		return ret, errors.New("Failed to parse page")
	}
	defer doc.Free()
	sectionXpath := xpath.Compile("//ul[contains(concat(' ', @class ,' '), ' slist ')]/li/a[1]")
	if sectionXpath != nil {
		defer sectionXpath.Free()
	}
	return ret, nil
}

func main() {
	fmt.Println(GetProxy())
}
