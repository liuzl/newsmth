package parser

import (
	"errors"
	"github.com/moovweb/gokogiri"
)

type DOMNode struct {
	Name string
	Node interface{}
	Item map[string]interface{}
}

func parseDOMRule(node interface{}, confRule ConfRule, pageUrl string) (interface{}, error) {

}

func Parse(page, url string, conf ParseConf) ([]ParsedItem, []ParsedTask, error) {
	var domList []*DOMNode
	var retUrls []ParsedTask
	var retItems []ParsedItem

	if conf == nil {
		return nil, nil, errors.New("parse conf is nil")
	}

	if len(page) == 0 {
		return nil, nil, errors.New("page len is 0")
	}

	root, err := gokogiri.ParseHtml([]byte(page))
	if err != nil {
		return nil, nil, err
	}
	defer root.Free()

	nodeItems := make(map[string]interface{})
	rootDOMNode := &DOMNode{Name: "root", Node: interface{}(root), Item: nodeItems}
	domList = append(domList, rootDOMNode)

	for {
		if len(domList) == 0 {
			break
		}
		domName := domList[0].Name
		domNode := domList[0].Node
		parentItems := domList[0].Item
		domList = domList[1:]
		var domConf []ConfRule
		var ok bool
		if domConf, ok = conf[domName]; !ok {
			continue
		}
	}
	return nil, nil, nil
}
