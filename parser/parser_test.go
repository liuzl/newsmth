package parser

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	page, _ := ioutil.ReadFile("./jrwen.html")
	pageUrl := "https://www.baidu.com/s?wd=%E6%96%87%E7%BB%A7%E8%8D%A3"
	urlRootConf := []ConfRule{
		ConfRule{Type: "dom", Name: "related", Xpath: "//div[@id='rs']//a"},
		ConfRule{Type: "string", Name: "title", Xpath: "//title/text()"},
	}
	urlDetailConf := []ConfRule{
		ConfRule{Type: "url", Name: "content_url", Xpath: "./@href"},
		ConfRule{Type: "string", Name: "url_text", Xpath: "./text()"},
	}

	urlConf := ParseConf{
		"root":    urlRootConf,
		"related": urlDetailConf,
	}

	retUrls, retItems, err := Parse(string(page), pageUrl, urlConf)
	if err != nil {
		t.Error(err)
	}
	t.Log("retUrls: ", retUrls)
	t.Log("retItems: ", retItems)
	jsonUrls, _ := json.Marshal(retUrls)
	jsonItems, _ := json.Marshal(retItems)
	t.Log("retUrls json: ", string(jsonUrls))
	t.Log("retItems json: ", string(jsonItems))
}
