package main

import (
	"fmt"
	"github.com/lestrrat/go-libxml2"
	"io/ioutil"
)

func main() {
	content, _ := ioutil.ReadFile("./section.html")
	doc, err := libxml2.ParseHTML(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer doc.Free()
	nodes, err := doc.Find("//tr[not(contains(td[2]/text(),'[二级目录]'))]/td[1]/a")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, node := range nodes.NodeList() {
		fmt.Println(node.TextContent())
		fmt.Println(node.String())
		fmt.Println(node.Find("./@href"))
	}
}
