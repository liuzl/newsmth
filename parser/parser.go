package parser

import (
	"errors"
	"fmt"
	"github.com/moovweb/gokogiri"
	"github.com/moovweb/gokogiri/html"
	"github.com/moovweb/gokogiri/xml"
	"github.com/moovweb/gokogiri/xpath"
	"strings"
)

type DOMNode struct {
	Name string
	Node interface{}
	Item map[string]interface{}
}

var (
	ErrInvalidXpath = errors.New("invalid xpath of node conf")
	ErrEmptyXpath   = errors.New("empty xpath of node conf")
	ErrInvalidType  = errors.New("invalid type of node conf")
	ErrEmptyType    = errors.New("empty type of node conf")
	ErrEmptyName    = errors.New("empty name of node conf")
)

func parseDOMRule(node interface{}, confRule ConfRule, pageUrl string) ([]interface{}, error) {
	if len(confRule.Xpath) == 0 {
		return nil, ErrEmptyXpath
	}
	if len(confRule.Type) == 0 {
		return nil, ErrEmptyType
	}
	nodeXpath := xpath.Compile(confRule.Xpath)
	if nodeXpath == nil {
		return nil, ErrInvalidXpath
	}
	defer nodeXpath.Free()

	var ret []interface{}
	var nodes []xml.Node
	var err error = nil
	switch node.(type) {
	case *html.HtmlDocument:
		nodes, err = node.(*html.HtmlDocument).Search(nodeXpath)
	case *xml.ElementNode:
		nodes, err = node.(*xml.ElementNode).Search(nodeXpath)
	default:
		err = errors.New("node type neither html.HtmlDocument nor xml.ElementNode")
	}
	if err != nil {
		return nil, errors.New(fmt.Sprintf("failed to xquery node, msg: %s", err))
	}
	for i, _ := range nodes {
		switch confRule.Type {
		case "dom":
			ret = append(ret, interface{}(nodes[i]))
		case "url":
			u, _ := MakeAbsoluteUrl(nodes[i].Content(), pageUrl)
			// TODO err
			ret = append(ret, interface{}(u))
		case "string":
			content := strings.TrimSpace(nodes[i].Content())
			if len(content) > 0 {
				ret = append(ret, interface{}(content))
			}
		}
	}
	return ret, err
}

func parseDOM(node interface{}, domConf []ConfRule, pageUrl string) ([]*DOMNode, []ParsedTask, map[string]interface{}, error) {
	var retDOMs []*DOMNode
	var retUrls []ParsedTask
	retItems := make(map[string]interface{})

	for _, conf := range domConf {
		if len(conf.Name) == 0 {
			return nil, nil, nil, ErrEmptyName
		}

		vals, err := parseDOMRule(node, conf, pageUrl)
		if err != nil {
			return nil, nil, nil, err
		}
		if len(conf.Regex) > 0 {
			var tmpVals []interface{}
			switch conf.Type {
			case "string":
				for _, v := range vals {
					res, err := ParseRegex(v.(string), conf.Regex)
					if err != nil {
						return nil, nil, nil, err
					}
					for _, tmpRes := range res {
						tmpVals = append(tmpVals, interface{}(tmpRes))
					}
				}
				vals = tmpVals
			case "url":
				for _, v := range vals {
					if MatchRegex(v.(string), conf.Regex) {
						tmpVals = append(tmpVals, v)
					}
				}
				vals = tmpVals
			}
		} //if has regex

		switch conf.Type {
		case "dom":
			for _, v := range vals {
				retDOMs = append(retDOMs, &DOMNode{
					Name: conf.Name,
					Node: interface{}(v),
					Item: retItems,
				})
			}
		case "url":
			for _, v := range vals {
				if u, ok := v.(string); ok {
					retUrls = append(retUrls, ParsedTask{TaskType: conf.Name, Url: u})
				}
			}
		case "string":
			if _, ok := retItems[conf.Name]; !ok {
				if len(vals) == 1 {
					retItems[conf.Name] = vals[0]
				} else if len(vals) > 1 {
					retItems[conf.Name] = interface{}(vals)
				}
			} else {
				switch retItems[conf.Name].(type) {
				case []interface{}:
					retItems[conf.Name] = append(retItems[conf.Name].([]interface{}), vals...)
				default:
					retItems[conf.Name] = append([]interface{}{retItems[conf.Name]}, vals...)
				}
				retItems[conf.Name] = interface{}(retItems[conf.Name])
			}
		}
	}
	return retDOMs, retUrls, retItems, nil
}

func Parse(page, pageUrl string, conf ParseConf) ([]ParsedTask, []ParsedItem, error) {
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
		if len(domList) == 0 { // no more dom to be processed
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
		parseDOMList, urlList, item, err := parseDOM(domNode, domConf, pageUrl)
		if err != nil {
			return nil, nil, err
		}
		domList = append(domList, parseDOMList...) // add more doms to be processed
		if urlList != nil {
			retUrls = append(retUrls, urlList...)
		}
		if item != nil {
			if _, ok = parentItems[domName]; !ok {
				parentItems[domName] = interface{}(item)
			} else {
				switch parentItems[domName].(type) {
				case []interface{}:
					parentItems[domName] = append(parentItems[domName].([]interface{}), interface{}(item))
				default:
					parentItems[domName] = []interface{}{parentItems[domName], interface{}(item)}
				}
				parentItems[domName] = interface{}(parentItems[domName])
			}
		}
	}
	if rootItems, ok := rootDOMNode.Item["root"]; ok {
		switch rootItems.(type) {
		case []interface{}:
			tmp, _ := rootItems.([]map[string]interface{})
			for _, v := range tmp {
				retItems = append(retItems, ParsedItem(v))
			}
		default:
			tmp, _ := rootItems.(map[string]interface{})
			retItems = append(retItems, ParsedItem(tmp))
		}
	}
	return retUrls, retItems, nil
}
