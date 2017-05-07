package parser

import (
	"reflect"
	"testing"
)

func TestParseRegex(t *testing.T) {
	content := "http://www.baidu.com"
	pattern0 := "http://www.(.+)"
	ret0, err := ParseRegex(content, pattern0)
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(ret0, []string{"baidu.com"}) {
		t.Error("pattern:", pattern0, ", result:", ret0)
	}

}
