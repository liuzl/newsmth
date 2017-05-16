package main

import (
	"fmt"
	"github.com/liuzl/newsmth/parser"
	"github.com/robertkrimen/otto"
	"io/ioutil"
	"log"
)

func main() {
	runtime := loadJsFromFile("test.js")
	if runtime == nil {
		log.Fatal("loadJsFrom File error")
	}

	r := new(parser.ConfRule)
	r.Type = "haha"
	v, err := runtime.ToValue(r)
	if err != nil {
		log.Fatal(err)
	}

	result, err := runtime.Call("process", nil, v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
	fmt.Println(r)
}

func loadJsFromFile(name string) *otto.Otto {
	js, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	runtime := otto.New()
	if _, err := runtime.Run(js); err != nil {
		log.Fatal(err)
	}
	return runtime
}
