# go-voyageai

go client for voyage ai

## Installation

```bash
go get github.com/conneroisu/go-voyageai
```

## Usage

```go
package main

import (
	"fmt"
	"log"

	"github.com/conneroisu/go-voyageai"
)

func main() {
	// create a new client
	client := voyageai.NewClient("YOUR_API_KEY")
	// create a new embedding request
	req := voyageai.EmbeddingsRequest{
		Model: string(voyageai.VoyageLarge2Instruct),
		Input: []string{"Hello, world!"},
	}
	// make the request
	res, err := client.Embeddings(req)
	if err != nil {
		log.Fatal(err)
	}
	// print the response is of type EmbeddingsResponse
	fmt.Println(res)
}
```

