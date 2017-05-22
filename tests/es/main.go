package main

import (
	"golang.org/x/net/context"
	elastic "gopkg.in/olivere/elastic.v5"
)

func main() {
	// Create a context
	ctx := context.Background()

	// Create a client
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	// Create an index
	_, err = client.CreateIndex("twitter").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
}
