package parser

import (
	"fmt"
	"regexp"
)

func ParseRegex(content, pattern string) ([]string, error) {
	fmt.Println(content, pattern)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	var ret []string
	res := re.FindAllStringSubmatch(content, -1)
	for i, _ := range res {
		switch {
		case len(res[i]) == 1:
			ret = append(ret, res[i][0])
		case len(res[i]) > 1:
			ret = append(ret, res[i][1:]...)
		}
	}
	fmt.Println(ret)
	return ret, nil
}

func MatchRegex(content, pattern string) bool {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}
	return re.MatchString(content)
}
