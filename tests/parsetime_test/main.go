package main

import (
	"fmt"
	"github.com/tkuchiki/parsetime"
	"log"
	"time"
)

func main() {
	fmt.Println("vim-go")
	p, err := parsetime.NewParseTime()
	if err != nil {
		log.Fatal(err)
	}
	strs := []string{"11:04:03 ", "2017-06-03"}
	for _, str := range strs {
		t, err := p.Parse(str)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(t, t.Unix(), time.Now().Unix(), time.Now())
	}
}
